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

type DMutexEtcd struct {
	c       *etcdclient.Client
	cess    *concurrency.Session
	timeout int //second
}

func NewDMutexEtcd(c *etcdclient.Client, timeout int) *DMutexEtcd {
	if timeout <= 0 {
		timeout = 10
	}
	m := DMutexEtcd{c: c, timeout: timeout}
	m.init()
	return &m
}

func (m *DMutexEtcd) init() error {
	if m.cess == nil {
		sess, err := concurrency.NewSession(m.c)
		if err != nil {
			zlog.LOG.Error("mutex init", zap.Error(err))
			return err
		}
		m.cess = sess
	}
	return nil
}

func (m *DMutexEtcd) Lock(pfx string) (*concurrency.Mutex, error) {
	stat := prome.NewStat("etcd_lock").MarkErr()
	defer stat.End()
	if err := m.init(); err != nil {
		return nil, err
	}

	mi := concurrency.NewMutex(m.cess, pfx)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(m.timeout))
	defer cancel()
	stat.MarkOk()
	return mi, mi.Lock(ctx)
}

func (m *DMutexEtcd) Unlock(mi *concurrency.Mutex) error {
	stat := prome.NewStat("etcd_unlock").MarkErr()
	defer stat.End()
	if mi != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(m.timeout))
		defer cancel()
		stat.MarkOk()
		return mi.Unlock(ctx)
	}
	return errors.New("mutex is nil")
}
