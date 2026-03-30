package dmutex

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/zlog"
	etcdclient "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.uber.org/zap"
)

// DMutexEtcd provides a distributed mutex backed by etcd. It is safe for
// concurrent use: the internal session is lazily created under mu.
type DMutexEtcd struct {
	mu      sync.Mutex
	c       *etcdclient.Client
	cess    *concurrency.Session
	timeout int // lock timeout in seconds
}

// NewDMutexEtcd returns a new DMutexEtcd using c as the etcd client.
// timeout is the per-operation deadline in seconds; values ≤ 0 default to 10.
func NewDMutexEtcd(c *etcdclient.Client, timeout int) *DMutexEtcd {
	if timeout <= 0 {
		timeout = 10
	}
	return &DMutexEtcd{c: c, timeout: timeout}
}

// init lazily creates the etcd session. It is called before every lock
// operation and is protected by mu to prevent concurrent session creation.
func (m *DMutexEtcd) init() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.cess == nil {
		sess, err := concurrency.NewSession(m.c)
		if err != nil {
			zlog.LOG.Error("failed to create etcd session", zap.Error(err))
			return err
		}
		m.cess = sess
	}
	return nil
}

// Lock acquires the distributed lock at pfx, blocking until it is available
// or the per-operation timeout elapses.
func (m *DMutexEtcd) Lock(pfx string) (*concurrency.Mutex, error) {
	stat := prome.NewStat("etcd_lock").MarkErr()
	defer stat.End()

	if err := m.init(); err != nil {
		return nil, err
	}

	mu := concurrency.NewMutex(m.cess, pfx)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(m.timeout))
	defer cancel()

	if err := mu.Lock(ctx); err != nil {
		zlog.LOG.Error("failed to acquire distributed lock",
			zap.String("prefix", pfx), zap.Error(err))
		return nil, err
	}

	stat.MarkOk()
	return mu, nil
}

// Unlock releases a distributed lock previously acquired by Lock or TryLock.
func (m *DMutexEtcd) Unlock(mu *concurrency.Mutex) error {
	stat := prome.NewStat("etcd_unlock").MarkErr()
	defer stat.End()

	if mu == nil {
		return errors.New("mutex is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(m.timeout))
	defer cancel()

	if err := mu.Unlock(ctx); err != nil {
		zlog.LOG.Error("failed to unlock distributed mutex", zap.Error(err))
		return err
	}

	stat.MarkOk()
	return nil
}

// Close closes the underlying etcd session and releases its lease.
// After Close, the DMutexEtcd must not be used.
func (m *DMutexEtcd) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.cess != nil {
		if err := m.cess.Close(); err != nil {
			zlog.LOG.Error("failed to close etcd session", zap.Error(err))
			return err
		}
		m.cess = nil
	}
	return nil
}

// TryLock tries to acquire the distributed lock at pfx, returning
// concurrency.ErrLocked immediately if another holder owns it.
func (m *DMutexEtcd) TryLock(pfx string) (*concurrency.Mutex, error) {
	stat := prome.NewStat("etcd_try_lock").MarkErr()
	defer stat.End()

	if err := m.init(); err != nil {
		return nil, err
	}

	mu := concurrency.NewMutex(m.cess, pfx)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(m.timeout))
	defer cancel()

	if err := mu.TryLock(ctx); err != nil {
		if !errors.Is(err, concurrency.ErrLocked) {
			zlog.LOG.Error("failed to try-acquire distributed lock",
				zap.String("prefix", pfx), zap.Error(err))
		}
		return nil, err
	}

	stat.MarkOk()
	return mu, nil
}
