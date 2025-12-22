package dmutex

import (
	"context"
	"errors"
	"time"

	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/zlog"
	etcdclient "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"go.uber.org/zap"
)

// DMutexEtcd provides a distributed mutex implementation using etcd
type DMutexEtcd struct {
	c       *etcdclient.Client
	cess    *concurrency.Session
	timeout int // Timeout in seconds
}

// NewDMutexEtcd creates a new distributed mutex instance with etcd backend
//
// @param c: etcd client instance
// @param timeout: lock timeout in seconds, defaults to 10 if <= 0
// @return: Pointer to initialized DMutexEtcd
func NewDMutexEtcd(c *etcdclient.Client, timeout int) *DMutexEtcd {
	if timeout <= 0 {
		timeout = 10
	}
	m := DMutexEtcd{c: c, timeout: timeout}
	return &m
}

// init initializes the etcd session if not already created
//
// @return: error if session creation fails
func (m *DMutexEtcd) init() error {
	if m.cess == nil {
		sess, err := concurrency.NewSession(m.c)
		if err != nil {
			zlog.LOG.Error("Failed to create etcd session for mutex", zap.Error(err))
			return err
		}
		m.cess = sess
	}
	return nil
}

// Lock acquires a distributed lock on the specified key prefix
//
// @param pfx: The key prefix to lock
// @return: *concurrency.Mutex and error if lock acquisition fails
func (m *DMutexEtcd) Lock(pfx string) (*concurrency.Mutex, error) {
	stat := prome.NewStat("etcd_lock").MarkErr()
	defer stat.End()

	if err := m.init(); err != nil {
		return nil, err
	}

	mu := concurrency.NewMutex(m.cess, pfx)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(m.timeout))
	defer cancel()

	err := mu.Lock(ctx)
	if err != nil {
		zlog.LOG.Error("Failed to acquire distributed lock",
			zap.String("prefix", pfx),
			zap.Error(err))
		return nil, err
	}

	stat.MarkOk()
	zlog.LOG.Debug("Successfully acquired distributed lock", zap.String("prefix", pfx))
	return mu, nil
}

// Unlock releases the distributed lock
//
// @param mu: The mutex to unlock
// @return: error if unlock operation fails
func (m *DMutexEtcd) Unlock(mu *concurrency.Mutex) error {
	stat := prome.NewStat("etcd_unlock").MarkErr()
	defer stat.End()

	if mu == nil {
		err := errors.New("mutex is nil")
		zlog.LOG.Error("Attempted to unlock nil mutex")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(m.timeout))
	defer cancel()

	err := mu.Unlock(ctx)
	if err != nil {
		zlog.LOG.Error("Failed to unlock distributed mutex", zap.Error(err))
		return err
	}

	stat.MarkOk()
	zlog.LOG.Debug("Successfully released distributed lock")
	return nil
}

// Close closes the etcd session and releases resources
//
// @return: error if session close fails
func (m *DMutexEtcd) Close() error {
	if m.cess != nil {
		err := m.cess.Close()
		if err != nil {
			zlog.LOG.Error("Failed to close etcd session", zap.Error(err))
			return err
		}
		m.cess = nil
		zlog.LOG.Debug("Successfully closed etcd session")
	}
	return nil
}

// TryLock attempts to acquire a lock with immediate return if lock is unavailable
//
// @param pfx: The key prefix to lock
// @return: *concurrency.Mutex and error if lock acquisition fails or lock is unavailable
func (m *DMutexEtcd) TryLock(pfx string) (*concurrency.Mutex, error) {
	stat := prome.NewStat("etcd_try_lock").MarkErr()
	defer stat.End()

	if err := m.init(); err != nil {
		return nil, err
	}

	mu := concurrency.NewMutex(m.cess, pfx)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(m.timeout))
	defer cancel()

	err := mu.TryLock(ctx)
	if err != nil {
		if errors.Is(err, concurrency.ErrLocked) {
			zlog.LOG.Debug("Distributed lock is already acquired by another client",
				zap.String("prefix", pfx))
		} else {
			zlog.LOG.Error("Failed to try acquire distributed lock",
				zap.String("prefix", pfx),
				zap.Error(err))
		}
		return nil, err
	}

	stat.MarkOk()
	zlog.LOG.Debug("Successfully acquired distributed lock with TryLock",
		zap.String("prefix", pfx))
	return mu, nil
}
