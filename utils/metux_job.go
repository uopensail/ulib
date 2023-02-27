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

type Register interface {
	// Register the registration.
	Register(ctx context.Context) error
	// Deregister the registration.
	Deregister(ctx context.Context) error
}

type MetuxJobUtil struct {
	reg        Register
	metux      *dmutex.DMutexEtcd
	prefixName string
	retry      int
}

func NewMetuxJobUtil(prefixName string, reg Register, etcdCli *etcdclient.Client, timeout, retry int) *MetuxJobUtil {
	metuxJob := MetuxJobUtil{}
	if timeout <= 0 {
		timeout = 10
	}
	if etcdCli != nil {
		metux := dmutex.NewDMutexEtcd(etcdCli, timeout)
		metuxJob.metux = metux
	}

	metuxJob.reg = reg
	metuxJob.prefixName = prefixName
	return &metuxJob
}

func (util *MetuxJobUtil) TryRun(job func()) {
	if job == nil {
		return
	}
	if util.metux == nil {
		job()
		return
	}
	for {
		lockKey := fmt.Sprintf("/lock_%s", util.prefixName)
		met, err := util.metux.Lock(lockKey)
		if err != nil {
			zlog.LOG.Warn("etcd.lock fail", zap.Error(err))
			if met != nil {
				util.metux.Unlock(met)
			}
			//抢锁失败等一会再抢
			time.Sleep(time.Second * time.Duration(10))
			continue
		} else {
			defer func() {
				if met != nil {
					err := util.metux.Unlock(met)
					if err != nil {
						zlog.LOG.Error("etcd.unlock fail", zap.Error(err), zap.String("key", lockKey))
					}
				}
			}()

			if util.reg != nil {
				myCtx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(20))
				defer cancel()
				err := util.reg.Deregister(myCtx)
				if err != nil {
					zlog.LOG.Warn("services Deregister", zap.Error(err))
				}
			}

			defer func() {
				if util.reg != nil {
					regCtx, regCancel := context.WithTimeout(context.Background(), time.Second*time.Duration(30))
					defer regCancel()
					err := util.reg.Register(regCtx)
					if err != nil {
						zlog.LOG.Warn("services Register", zap.Error(err))
					}
				}
			}()

			job()
			return
		}

	}
}
