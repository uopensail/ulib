package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/uopensail/ulib/dmutex"
	"github.com/uopensail/ulib/zlog"
	etcdclient "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

// Register abstracts service registration and deregistration in a service
// discovery system (e.g. etcd-based service mesh).
type Register interface {
	// Register publishes the service instance to the registry.
	Register(ctx context.Context) error
	// Deregister removes the service instance from the registry.
	Deregister(ctx context.Context) error
}

// MetuxJobUtil runs a job exclusively across a cluster by acquiring a
// distributed etcd mutex before execution. If no etcd client is provided the
// job runs locally without any distributed locking.
type MetuxJobUtil struct {
	reg        Register
	metux      *dmutex.DMutexEtcd
	prefixName string
}

// NewMetuxJobUtil constructs a MetuxJobUtil.
//
//   - prefixName: key prefix used for the distributed lock (e.g. "reload").
//   - reg: optional service-registry handle; may be nil.
//   - etcdCli: etcd client used for locking; pass nil to disable distributed locking.
//   - timeout: lock acquisition timeout in seconds (clamped to 10 if ≤ 0).
func NewMetuxJobUtil(prefixName string, reg Register, etcdCli *etcdclient.Client, timeout int) *MetuxJobUtil {
	u := &MetuxJobUtil{
		reg:        reg,
		prefixName: prefixName,
	}
	if etcdCli != nil {
		if timeout <= 0 {
			timeout = 10
		}
		u.metux = dmutex.NewDMutexEtcd(etcdCli, timeout)
	}
	return u
}

// TryRun executes job once the distributed lock is held. If locking fails it
// retries after 10 s. When no etcd client was configured the job is executed
// immediately without locking.
//
// Service deregistration is performed before the job and re-registration
// after, so that traffic is not routed to this node while the job runs.
func (u *MetuxJobUtil) TryRun(job func()) {
	if job == nil {
		return
	}
	if u.metux == nil {
		// No distributed locking configured — run directly.
		job()
		return
	}

	lockKey := fmt.Sprintf("/lock_%s", u.prefixName)
	for {
		if u.tryRunOnce(lockKey, job) {
			return
		}
		// Lock acquisition failed; wait before retrying.
		time.Sleep(10 * time.Second)
	}
}

// tryRunOnce attempts to acquire the lock, run the job, then release the lock.
// It returns true when the job was executed successfully, false when the lock
// could not be acquired so the caller should retry.
//
// Using a separate function ensures that deferred calls (unlock, re-register)
// fire immediately after the job completes rather than at the end of the
// outer for-loop in TryRun.
func (u *MetuxJobUtil) tryRunOnce(lockKey string, job func()) (done bool) {
	mu, err := u.metux.Lock(lockKey)
	if err != nil {
		zlog.LOG.Warn("etcd.lock fail", zap.Error(err))
		if mu != nil {
			// Lock returned an error but also a non-nil mutex — unlock defensively.
			if uerr := u.metux.Unlock(mu); uerr != nil {
				zlog.LOG.Error("etcd.unlock fail after failed lock", zap.Error(uerr))
			}
		}
		return false
	}

	// Ensure the lock is always released when this function returns.
	defer func() {
		if uerr := u.metux.Unlock(mu); uerr != nil {
			zlog.LOG.Error("etcd.unlock fail", zap.Error(uerr), zap.String("key", lockKey))
		}
	}()

	// Deregister before running the job so no traffic is directed here.
	if u.reg != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		if derr := u.reg.Deregister(ctx); derr != nil {
			zlog.LOG.Warn("services Deregister failed", zap.Error(derr))
		}
	}

	// Re-register after the job finishes regardless of outcome.
	if u.reg != nil {
		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if rerr := u.reg.Register(ctx); rerr != nil {
				zlog.LOG.Warn("services Register failed", zap.Error(rerr))
			}
		}()
	}

	job()
	return true
}
