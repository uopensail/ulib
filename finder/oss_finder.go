package finder

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/utils"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
	"strings"
	"time"
)

type OSSFinder struct {
	Config *commonconfig.FinderConfig
}

func NewOSSFinder(conf *commonconfig.FinderConfig) *OSSFinder {
	return &OSSFinder{
		Config: conf,
	}
}

func (finder *OSSFinder) getBucket(path string) (string, error) {
	if !strings.HasPrefix(path, "oss://") {
		return "", fmt.Errorf("%s is not oss path", path)
	}
	return strings.Split(path[6:], "/")[0], nil
}

func (finder *OSSFinder) getPath(path string) (string, error) {
	if !strings.HasPrefix(path, "oss://") {
		return "", fmt.Errorf("%s is not oss path", path)
	}
	item := strings.Split(path[6:], "/")[0]
	return path[len("oss://"+item+"/"):], nil
}

func (finder *OSSFinder) ListDir(dir string) ([]string, error) {
	stat := prome.NewStat("OSSFinder.ListDir")
	defer stat.End()
	bucket, err := finder.getBucket(dir)
	if err != nil {
		stat.MarkErr()
		return nil, err
	}
	path, err := finder.getPath(dir)
	if err != nil {
		stat.MarkErr()
		return nil, err
	}
	timeout := 60
	if finder.Config.Timeout > 0 {
		timeout = finder.Config.Timeout
	}

	client, err := oss.New(
		finder.Config.Endpoint,
		finder.Config.AccessKey,
		finder.Config.SecretKey,
		oss.Timeout(int64(timeout), int64(timeout)),
	)
	if err != nil {
		stat.MarkErr()
		return nil, err
	}
	ossBucket, err := client.Bucket(bucket)
	if err != nil {
		stat.MarkErr()
		return nil, err
	}
	resp, err := ossBucket.ListObjects(oss.Prefix(path))

	if err != nil {
		stat.MarkErr()
		return nil, err
	}

	ret := make([]string, 0, len(resp.Objects))
	for _, value := range resp.Objects {
		ret = append(ret, "oss://"+bucket+"/"+value.Key)
	}
	return ret, nil
}

func (finder *OSSFinder) Download(src, dst string) (int64, error) {
	stat := prome.NewStat("OSSFinder.Download")
	defer stat.End()
	bucket, err := finder.getBucket(src)
	if err != nil {
		stat.MarkErr()
		return 0, err
	}
	path, err := finder.getPath(src)
	if err != nil {
		stat.MarkErr()
		return 0, err
	}
	timeout := 60
	if finder.Config.Timeout > 0 {
		timeout = finder.Config.Timeout
	}

	client, err := oss.New(
		finder.Config.Endpoint,
		finder.Config.AccessKey,
		finder.Config.SecretKey,
		oss.Timeout(int64(timeout), int64(timeout)),
	)
	if err != nil {
		stat.MarkErr()
		return 0, err
	}
	ossBucket, err := client.Bucket(bucket)
	if err != nil {
		stat.MarkErr()
		return 0, err
	}
	err = ossBucket.GetObjectToFile(path, dst)

	if err != nil {
		stat.MarkErr()
		return 0, err
	}

	head, err := ossBucket.GetObjectMeta(path)

	if err != nil {
		return 0, err
	}
	length := utils.String2Int64(head.Get("Content-Length"))
	zlog.LOG.Info("OSS Finder", zap.String("path", src), zap.Int64("length", length))
	return length, nil
}

func (finder *OSSFinder) GetUpdateTime(filepath string) int64 {
	stat := prome.NewStat("OSSFinder.GetUpdateTime")
	defer stat.End()
	bucket, err := finder.getBucket(filepath)
	if err != nil {
		stat.MarkErr()
		return -1
	}
	path, err := finder.getPath(filepath)
	if err != nil {
		stat.MarkErr()
		return -1
	}
	timeout := 60
	if finder.Config.Timeout > 0 {
		timeout = finder.Config.Timeout
	}

	client, err := oss.New(
		finder.Config.Endpoint,
		finder.Config.AccessKey,
		finder.Config.SecretKey,
		oss.Timeout(int64(timeout), int64(timeout)),
	)
	if err != nil {
		stat.MarkErr()
		return -1
	}
	ossBucket, err := client.Bucket(bucket)
	if err != nil {
		stat.MarkErr()
		return -1
	}
	head, err := ossBucket.GetObjectMeta(path)

	if err != nil {
		stat.MarkErr()
		return -1
	}

	lastModified := head.Get("Last-Modified")
	ts, err := time.Parse(time.RFC1123, lastModified)
	if err != nil {
		stat.MarkErr()
		return -1
	}
	return ts.Unix()
}

func (finder *OSSFinder) GetETag(filepath string) string {
	stat := prome.NewStat("OSSFinder.GetETag")
	defer stat.End()
	bucket, err := finder.getBucket(filepath)
	if err != nil {
		stat.MarkErr()
		return ""
	}
	path, err := finder.getPath(filepath)
	if err != nil {
		stat.MarkErr()
		return ""
	}
	timeout := 60
	if finder.Config.Timeout > 0 {
		timeout = finder.Config.Timeout
	}

	client, err := oss.New(
		finder.Config.Endpoint,
		finder.Config.AccessKey,
		finder.Config.SecretKey,
		oss.Timeout(int64(timeout), int64(timeout)),
	)
	if err != nil {
		stat.MarkErr()
		return ""
	}
	ossBucket, err := client.Bucket(bucket)
	if err != nil {
		stat.MarkErr()
		return ""
	}
	head, err := ossBucket.GetObjectMeta(path)

	if err != nil {
		stat.MarkErr()
		return ""
	}

	return head.Get("ETag")
}

func (finder *OSSFinder) GetLocalETag(filepath string) string {
	stat := prome.NewStat("OSSFinder.GetLocalETag")
	defer stat.End()
	bucket, err := finder.getBucket(filepath)
	if err != nil {
		stat.MarkErr()
		return ""
	}
	path, err := finder.getPath(filepath)
	if err != nil {
		stat.MarkErr()
		return ""
	}
	timeout := 60
	if finder.Config.Timeout > 0 {
		timeout = finder.Config.Timeout
	}

	client, err := oss.New(
		finder.Config.Endpoint,
		finder.Config.AccessKey,
		finder.Config.SecretKey,
		oss.Timeout(int64(timeout), int64(timeout)),
	)
	if err != nil {
		stat.MarkErr()
		return ""
	}
	ossBucket, err := client.Bucket(bucket)
	if err != nil {
		stat.MarkErr()
		return ""
	}
	head, err := ossBucket.GetObjectMeta(path)

	if err != nil {
		stat.MarkErr()
		return ""
	}

	return head.Get("ETag")
}
