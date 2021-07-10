package loader

import (
	"fmt"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/finder"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/utils"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Status int

const (
	DownloadOK Status = iota + 1
	RemoteFileError
	LocalFileIsNewest
	DownloadFileError
)

type ITable interface {
}

type CreateFunc func(interface{}) ITable

type Job struct {
	Finder         finder.IFinder
	Interval       int
	LastUpdateTime int64
}

type Manager struct {
	Locker   *sync.RWMutex
	TableMap map[string]ITable
	JobMap   map[string]*Job
}

func init() {
	ManagerImp = &Manager{
		Locker:   new(sync.RWMutex),
		TableMap: make(map[string]ITable),
		JobMap:   make(map[string]*Job),
	}
}

func Register(key string, cfg *commonconfig.DownloaderConfig, factory CreateFunc, params interface{}) bool {
	interval := 30
	if cfg.Interval > interval {
		interval = cfg.Interval
	}
	record := &Job{
		Interval:       interval,
		Finder:         finder.GetFinder(&cfg.FinderConfig),
		LastUpdateTime: 0,
	}

	status, remoteUpdateTime := download(record.Finder, cfg.SourcePath, cfg.LocalPath)
	switch status {
	case DownloadOK, LocalFileIsNewest:
		record.LastUpdateTime = remoteUpdateTime
	case DownloadFileError, RemoteFileError:
		return false
	}
	//加载table
	table := factory(params)
	ManagerImp.Locker.Lock()
	ManagerImp.TableMap[key] = table
	ManagerImp.JobMap[key] = record
	ManagerImp.Locker.Unlock()

	go func(r *Job) {
		ticker := time.NewTicker(time.Second * time.Duration(r.Interval))
		defer ticker.Stop()
		for {
			<-ticker.C
			realUpdateTime := r.Finder.GetUpdateTime(cfg.SourcePath)
			if realUpdateTime < r.LastUpdateTime {
				continue
			}

			status, remoteUpdateTime = download(record.Finder, cfg.SourcePath, cfg.LocalPath)
			switch status {
			case DownloadOK:
				record.LastUpdateTime = remoteUpdateTime
			case DownloadFileError, RemoteFileError:
				prome.NewStat(fmt.Sprintf("Loader.%s", key)).MarkErr().End()
				continue
			case LocalFileIsNewest:
				record.LastUpdateTime = remoteUpdateTime
				prome.NewStat(fmt.Sprintf("Loader.%s", key)).MarkOk().End()
				continue
			}

			table = factory(params)
			ManagerImp.Locker.Lock()
			ManagerImp.TableMap[key] = table
			ManagerImp.Locker.Unlock()
			r.LastUpdateTime = realUpdateTime
			prome.NewStat(fmt.Sprintf("Loader.%s", key)).MarkOk().End()
			zlog.LOG.Info("Loader", zap.String("source", cfg.SourcePath), zap.String("local", cfg.LocalPath))
		}
	}(record)
	return true
}

func GetTable(key string) ITable {
	ManagerImp.Locker.RLocker()
	defer ManagerImp.Locker.RUnlock()
	if v, ok := ManagerImp.TableMap[key]; ok {
		return v
	}
	return nil
}

func download(finder finder.IFinder, remotePath, localPath string) (Status, int64) {
	localUpdateTime := utils.GetFileModifyTime(localPath)
	remoteUpdateTime := finder.GetUpdateTime(remotePath)
	//远程找不到文件注册出错
	if remoteUpdateTime == -1 {
		zlog.LOG.Error(fmt.Sprintf("Get %s Modify Time error", remotePath))
		return RemoteFileError, -1
	}

	//远程的文件更新了，或者本地没有文件
	if remoteUpdateTime > localUpdateTime {
		//下载文件
		size, err := finder.Download(remotePath, localPath)
		if err != nil || size == 0 {
			zlog.LOG.Error(fmt.Sprintf("DownLoader %s error", remotePath))
			return DownloadFileError, -1
		}
		return DownloadOK, remoteUpdateTime
	}
	return LocalFileIsNewest, remoteUpdateTime
}

var ManagerImp *Manager
