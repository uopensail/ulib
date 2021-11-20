package loader

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/finder"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/utils"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
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

type CreateFunc func(dataPath string, params interface{}) ITable

type ReleaseFunc func(ITable, interface{})

type Job struct {
	Key            string
	Finder         finder.IFinder
	Interval       int
	IterCount      int64
	DownloadConfig *commonconfig.DownloaderConfig
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

//Register 注册任务
func Register(key string, cfg *commonconfig.DownloaderConfig,
	factory CreateFunc, release ReleaseFunc,
	createParams, releaseParams interface{}) bool {
	interval := 30
	if cfg.Interval > 0 {
		interval = cfg.Interval
	}

	myFinder := finder.GetFinder(&cfg.FinderConfig)
	status, iterCount := tryDownloadIfNeed(myFinder, cfg.SourcePath, cfg.LocalPath)

	switch status {
	case DownloadFileError, RemoteFileError, WriteLocalETagError:
		return false
	}
	record := &Job{
		Key:            key,
		Interval:       interval,
		Finder:         myFinder,
		IterCount:      iterCount,
		DownloadConfig: cfg,
	}
	//加载table
	table := factory(fmt.Sprintf("%s.%d", record.DownloadConfig.LocalPath, iterCount), createParams)
	if table == nil {
		return false
	}
	ManagerImp.Locker.Lock()
	ManagerImp.TableMap[key] = table
	ManagerImp.JobMap[key] = record
	ManagerImp.Locker.Unlock()

	go func(r *Job) {
		ticker := time.NewTicker(time.Second * time.Duration(r.Interval))
		defer ticker.Stop()
		for {
			<-ticker.C
			status, iterCount = tryDownloadIfNeed(r.Finder, r.DownloadConfig.SourcePath, r.DownloadConfig.LocalPath)
			switch status {
			case DownloadOK:
				record.IterCount = iterCount
			case DownloadFileError, RemoteFileError, WriteLocalETagError:
				prome.NewStat(fmt.Sprintf("Loader.%s", key)).MarkErr().End()
				continue
			case LocalFileIsNewest:
				prome.NewStat(fmt.Sprintf("Loader.%s", r.Key)).MarkOk().End()
				//zlog.LOG.Info(fmt.Sprintf("Loader: %s do not need reload", r.Key),
				//	zap.String("source", r.DownloadConfig.SourcePath),
				//	zap.String("local", fmt.Sprintf("%s.%d", r.DownloadConfig.LocalPath, r.IterCount)),
				//)
				continue
			}

			table = factory(fmt.Sprintf("%s.%d", r.DownloadConfig.LocalPath, r.IterCount), createParams)
			if table == nil {
				zlog.LOG.Error("create table nil")
				continue
			}
			ManagerImp.Locker.Lock()
			old, ok := ManagerImp.TableMap[key]
			ManagerImp.TableMap[key] = table
			ManagerImp.Locker.Unlock()

			if ok {
				//延迟释放
				go func(obj ITable, p interface{}, count int64, localPath string) {
					time.Sleep(time.Second)
					if release != nil && obj != nil {
						release(obj, p)
					}
					dFilePath := fmt.Sprintf("%s.%d", localPath, count-1)
					metaPath := fmt.Sprintf("%s.%d_meta", localPath, count-1)
					zlog.LOG.Info("loader.release.Remove", zap.String("path", dFilePath))
					zlog.LOG.Info("loader.release.Remove.meta", zap.String("path", metaPath))
					os.Remove(dFilePath)
					os.Remove(metaPath)
				}(old, releaseParams, iterCount, r.DownloadConfig.LocalPath)
			}

			prome.NewStat(fmt.Sprintf("Loader.%s", r.Key)).MarkOk().End()
			//zlog.LOG.Info(fmt.Sprintf("Loader: %s reload", r.Key),
			//	zap.String("source", r.DownloadConfig.SourcePath),
			//	zap.String("local", fmt.Sprintf("%s.%d", r.DownloadConfig.LocalPath, r.IterCount)),
			//)
		}
	}(record)
	return true
}

func GetTable(key string) ITable {
	ManagerImp.Locker.RLock()
	defer ManagerImp.Locker.RUnlock()
	if v, ok := ManagerImp.TableMap[key]; ok {
		return v
	}
	return nil
}

func readLocalETag(localPath string) string {
	f, err := os.OpenFile(localPath, os.O_RDONLY, 0600)
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
	f, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(etag))
	return err
}

func getNewestFile(filename string) (string, int64) {
	index := -1
	files, _ := utils.ListDir(filepath.Dir(filename))
	for i := 0; i < len(files); i++ {
		if strings.HasPrefix(files[i], filename+".") {
			v, err := strconv.Atoi(files[i][len(filename)+1:])
			if err != nil {
				continue
			}
			if v > index {
				index = v
			}
		}
	}
	if index == -1 {
		return "", -1
	}
	return fmt.Sprintf("%s.%d", filename, index), int64(index)
}

//tryDownloadIfNeed 下载文件
func tryDownloadIfNeed(finder finder.IFinder, remotePath, localPath string) (Status, int64) {
	remoteETag := finder.GetETag(remotePath)
	if len(remoteETag) == 0 {
		zlog.LOG.Error("GetETag", zap.String("remote_path", remotePath))
		return RemoteFileError, -1
	}

	localFileName, iterCount := getNewestFile(localPath)

	//本地有文件
	if len(localFileName) > 0 {
		localETag := readLocalETag(localFileName + "_meta")
		//etag相等，或者远程etag出错，就不需要下载
		if remoteETag == localETag {
			return LocalFileIsNewest, iterCount
		}
	}

	iterCount++

	//需要下载文件
	size, err := finder.Download(remotePath, fmt.Sprintf("%s.%d", localPath, iterCount))
	if err != nil || size == 0 {
		zlog.LOG.Error("Download", zap.String("remote_path", remotePath), zap.Error(err))
		return DownloadFileError, -1
	}
	err = writeLocalETag(fmt.Sprintf("%s.%d_meta", localPath, iterCount), remoteETag)
	if err != nil {
		zlog.LOG.Error("writeLocalETag", zap.String("local_path", localPath), zap.Error(err))
		return WriteLocalETagError, -1
	}
	return DownloadOK, iterCount
}

var ManagerImp *Manager
