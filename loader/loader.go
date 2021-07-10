package loader

import (
	"fmt"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/finder"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
	"sync"
	"time"
)

type ITable interface {
	Load(string) (int, error)
}

type LoadFunc func(cfg *commonconfig.DownloaderConfig) ITable

type JobRecord struct {
	Status   bool
	Config   *commonconfig.DownloaderConfig
	Finder   finder.IFinder
	Func     LoadFunc
	Interval int
}
type Manager struct {
	Locker   *sync.RWMutex
	TableMap map[string]ITable
	JobMap   map[string]*JobRecord
}

var ManagerImp *Manager

func init() {
	ManagerImp = &Manager{
		Locker:   new(sync.RWMutex),
		TableMap: make(map[string]ITable),
		JobMap:   make(map[string]*JobRecord),
	}
}

func Register(key string, cfg *commonconfig.DownloaderConfig, foo LoadFunc) {
	interval := 60
	if cfg.Interval > interval {
		interval = cfg.Interval
	}
	record := &JobRecord{
		Status:   true,
		Config:   cfg,
		Interval: interval,
		Finder:   finder.GetFinder(&cfg.FinderConfig),
		Func:     foo,
	}
	//添加一个定时任务
	ManagerImp.Locker.Lock()
	ManagerImp.JobMap[key] = record
	ManagerImp.Locker.Unlock()

	go func(r *JobRecord) {
		ticker := time.NewTicker(time.Second * time.Duration(r.Interval))
		defer ticker.Stop()
		lastUpdateTime := int64(0)
		for {
			<-ticker.C
			if !r.Status {
				break
			}
			realUpdateTime := r.Finder.GetUpdateTime(cfg.Source)
			if realUpdateTime <= lastUpdateTime {
				continue
			}
			stat := prome.NewStat(fmt.Sprintf("Loader.%s", key))
			table := r.Func(cfg)
			ManagerImp.Locker.Lock()
			ManagerImp.TableMap[key] = table
			ManagerImp.Locker.Unlock()
			stat.End()
			zlog.LOG.Info("Loader", zap.String("source", cfg.Source), zap.String("local", cfg.Local))
		}
	}(record)
}

func Remove(key string) {
	ManagerImp.Locker.Lock()
	defer ManagerImp.Locker.Unlock()
	if v, ok := ManagerImp.JobMap[key]; ok {
		v.Status = false
		delete(ManagerImp.JobMap, key)
	}
}

func GetTable(key string) ITable {
	ManagerImp.Locker.RLocker()
	defer ManagerImp.Locker.RUnlock()
	if v, ok := ManagerImp.TableMap[key]; ok {
		return v
	}
	return nil
}
