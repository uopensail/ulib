package loader

import (
	"fmt"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/finder"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type Status int

const (
	DownloadOK Status = iota + 1
	RemoteFileError
	LocalFileIsNewest
	DownloadFileError
	WriteLocalETagError
)

type ITable interface{}

type CreateFunc func(interface{}) ITable

type ReleaseFunc func(ITable, interface{})

type Job struct {
	Key            string
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

func Register(key string, cfg *commonconfig.DownloaderConfig, factory CreateFunc,
	release ReleaseFunc, createParams, releaseParams interface{}) bool {
	interval := 30
	if cfg.Interval > 0 {
		interval = cfg.Interval
	}
	record := &Job{
		Key:            key,
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
	table := factory(createParams)
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
				prome.NewStat(fmt.Sprintf("Loader.%s", r.Key)).MarkOk().End()
				zlog.LOG.Info(fmt.Sprintf("Loader: %s do not need reload", r.Key),
					zap.String("source", cfg.SourcePath),
					zap.String("local", cfg.LocalPath),
					zap.Int64("updateTime", remoteUpdateTime),
				)
				continue
			}

			status, remoteUpdateTime = download(record.Finder, cfg.SourcePath, cfg.LocalPath)
			switch status {
			case DownloadOK:
				record.LastUpdateTime = remoteUpdateTime
			case DownloadFileError, RemoteFileError, WriteLocalETagError:
				prome.NewStat(fmt.Sprintf("Loader.%s", key)).MarkErr().End()
				continue
			case LocalFileIsNewest:
				record.LastUpdateTime = remoteUpdateTime
				prome.NewStat(fmt.Sprintf("Loader.%s", r.Key)).MarkOk().End()
				zlog.LOG.Info(fmt.Sprintf("Loader: %s do not need reload", r.Key),
					zap.String("source", cfg.SourcePath),
					zap.String("local", cfg.LocalPath),
					zap.Int64("updateTime", remoteUpdateTime),
				)
				continue
			}

			table = factory(createParams)
			ManagerImp.Locker.Lock()
			old, ok := ManagerImp.TableMap[key]
			ManagerImp.TableMap[key] = table
			ManagerImp.Locker.Unlock()

			if ok {
				//延迟释放
				go func(obj ITable, p interface{}) {
					time.Sleep(time.Second)
					release(obj, p)
				}(old, releaseParams)
			}

			r.LastUpdateTime = realUpdateTime
			prome.NewStat(fmt.Sprintf("Loader.%s", r.Key)).MarkOk().End()
			zlog.LOG.Info(fmt.Sprintf("Loader: %s reload", r.Key),
				zap.String("source", cfg.SourcePath),
				zap.String("local", cfg.LocalPath),
				zap.Int64("updateTime", remoteUpdateTime),
			)
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

func readLocalETag(localPath string) string {
	f, err := os.OpenFile(fmt.Sprintf("%s_success", localPath), os.O_RDONLY, 0600)
	if err != nil {
		return ""
	}
	defer f.Close()

	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		return ""
	}
	return string(contentByte)
}

func writeLocalETag(localPath string, etag string) error {
	f, err := os.Create(fmt.Sprintf("%s_success", localPath))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(etag))
	return err
}

func download(finder finder.IFinder, remotePath, localPath string) (Status, int64) {
	localETag := readLocalETag(localPath)
	remoteETag := finder.GetETag(remotePath)
	remoteUpdateTime := finder.GetUpdateTime(remotePath)
	//远程找不到文件注册出错
	if len(remoteETag) == 0 {
		zlog.LOG.Error(fmt.Sprintf("Get %s ETag error", remotePath))
		return RemoteFileError, -1
	}

	//不需要下载
	if localETag == remoteETag {
		return LocalFileIsNewest, remoteUpdateTime
	}

	//下载文件
	size, err := finder.Download(remotePath, localPath)
	if err != nil || size == 0 {
		zlog.LOG.Error(fmt.Sprintf("DownLoader %s error", remotePath))
		return DownloadFileError, -1
	}
	err = writeLocalETag(localPath, remoteETag)
	if err != nil {
		return WriteLocalETagError, -1
	}
	return DownloadOK, remoteUpdateTime
}

var ManagerImp *Manager
