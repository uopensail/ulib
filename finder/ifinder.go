package finder

import (
	"fmt"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
	"os"
	"sync"
	"time"
)

type IFinder interface {
	ListDir(dir string) ([]string, error)
	Download(src, dst string) (int64, error)
	GetUpdateTime(filepath string) int64
}

type Job struct {
	Finder   IFinder
	Interval int
}

type Meta struct {
	Path       string
	ModifyTime int64
}

func GetFinder(cfg *commonconfig.FinderConfig) IFinder {
	switch cfg.Type {
	case "s3":
		return NewS3Finder(cfg)
	case "oss":
		return NewOSSFinder(cfg)
	default:
		return nil
	}
}

type Manager struct {
	Locker  *sync.RWMutex
	FileMap map[string]*Meta
	JobMap  map[string]*Job
}

func init() {
	ManagerImp = &Manager{
		Locker:  new(sync.RWMutex),
		FileMap: make(map[string]*Meta),
		JobMap:  make(map[string]*Job),
	}
}

func Register(key string, cfg *commonconfig.DownloaderConfig) bool {
	interval := 30
	if cfg.Interval > interval {
		interval = cfg.Interval
	}
	record := &Job{
		Interval: interval,
		Finder:   GetFinder(&cfg.FinderConfig),
	}

	//检查本地文件的修改时间与远程文件的修改时间
	lastLocalModifyTime := getModifyTime(cfg.LocalPath)
	lastRemoteModifyTime := record.Finder.GetUpdateTime(cfg.SourcePath)
	if lastRemoteModifyTime == -1 {
		prome.NewStat(fmt.Sprintf("Finder.%s", key)).MarkErr().End()
		return false
	}

	if lastRemoteModifyTime > lastLocalModifyTime {
		//下载文件
		_, err := record.Finder.Download(cfg.SourcePath, cfg.LocalPath)
		if err != nil {
			prome.NewStat(fmt.Sprintf("Finder.%s", key)).MarkErr().End()
			return false
		}
	}
	ManagerImp.Locker.Lock()
	ManagerImp.FileMap[key] = &Meta{
		Path:       cfg.LocalPath,
		ModifyTime: lastRemoteModifyTime,
	}
	ManagerImp.JobMap[key] = record
	ManagerImp.Locker.Unlock()

	go func(r *Job) {
		ticker := time.NewTicker(time.Second * time.Duration(r.Interval))
		defer ticker.Stop()
		for {
			<-ticker.C
			realUpdateTime := r.Finder.GetUpdateTime(cfg.SourcePath)
			if realUpdateTime < Get(key).ModifyTime {
				continue
			}
			_, err := r.Finder.Download(cfg.SourcePath, cfg.LocalPath)
			if err != nil {
				prome.NewStat(fmt.Sprintf("Finder.%s", key)).MarkErr().End()
				continue
			}
			ManagerImp.Locker.Lock()
			ManagerImp.FileMap[key] = &Meta{
				Path:       cfg.LocalPath,
				ModifyTime: realUpdateTime,
			}
			ManagerImp.Locker.Unlock()
			prome.NewStat(fmt.Sprintf("Finder.%s", key)).MarkOk().End()
			zlog.LOG.Info("Finder", zap.String("source", cfg.SourcePath), zap.String("local", cfg.LocalPath))
		}
	}(record)
	return true
}

func Get(key string) *Meta {
	ManagerImp.Locker.RLocker()
	defer ManagerImp.Locker.RUnlock()
	if v, ok := ManagerImp.FileMap[key]; ok {
		return v
	}
	return nil
}

func getModifyTime(fileName string) int64 {
	info, err := os.Stat(fileName)
	if err != nil {
		return -1
	}
	return info.ModTime().Unix()
}

var ManagerImp *Manager
