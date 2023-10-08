package utils

import (
	"context"
	"fmt"
	"sync/atomic"
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

type JoberStatus uint32

const (
	NormalJobStatus JoberStatus = 1
	DoingJobStatus  JoberStatus = 2
)

type MetuxJobUtil struct {
	reg        Register
	metux      *dmutex.DMutexEtcd
	prefixName string
	retry      int
	jobStatus  uint32
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
	metuxJob.jobStatus = uint32(NormalJobStatus)
	return &metuxJob
}

func (util *MetuxJobUtil) JobStatus() JoberStatus {
	return JoberStatus(atomic.LoadUint32(&util.jobStatus))
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
			// sleep 10s and try again
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

			atomic.StoreUint32(&(util.jobStatus), uint32(DoingJobStatus))
			defer func() {
				atomic.StoreUint32(&util.jobStatus, uint32(NormalJobStatus))
			}()

			job()
			return
		}

	}
}
