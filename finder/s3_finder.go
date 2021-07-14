package finder

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

type S3Finder struct {
	Config  *commonconfig.FinderConfig
	Session *session.Session
}

func NewS3Finder(conf *commonconfig.FinderConfig) *S3Finder {
	finder := &S3Finder{
		Config: conf,
	}
	awsConf := aws.Config{S3ForcePathStyle: aws.Bool(false)}
	if len(conf.AccessKey) > 0 && len(conf.SecretKey) > 0 {
		awsConf.Credentials = credentials.NewStaticCredentials(conf.AccessKey, conf.SecretKey, "")
		awsConf.DisableSSL = aws.Bool(true)
	}

	if len(conf.Endpoint) > 0 {
		awsConf.Endpoint = aws.String(conf.Endpoint)
	}

	if len(conf.Region) > 0 {
		awsConf.Region = aws.String(conf.Region)
	} else {
		awsConf.Region = aws.String(endpoints.EuWest1RegionID)
	}
	sess, err := session.NewSession(&awsConf)
	if err != nil {
		panic(err)
	}
	finder.Session = sess
	return finder
}

func (finder *S3Finder) getBucket(path string) (string, error) {
	if !strings.HasPrefix(path, "s3://") {
		return "", fmt.Errorf("%s is not s3 path", path)
	}
	return strings.Split(path[5:], "/")[0], nil
}

func (finder *S3Finder) getPath(path string) (string, error) {
	if !strings.HasPrefix(path, "s3://") {
		return "", fmt.Errorf("%s is not s3 path", path)
	}
	item := strings.Split(path[5:], "/")[0]
	return path[len("s3://"+item+"/"):], nil
}

func (finder *S3Finder) ListDir(dir string) ([]string, error) {
	stat := prome.NewStat("S3Finder.ListDir")
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
	client := s3.New(finder.Session)
	resp, err := client.ListObjects(
		&s3.ListObjectsInput{
			Bucket:    aws.String(bucket),
			Prefix:    aws.String(path),
			Delimiter: aws.String("/"),
		})
	if err != nil {
		stat.MarkErr()
		return nil, err
	}

	ret := make([]string, 0, len(resp.Contents))
	for _, value := range resp.Contents {
		ret = append(ret, "s3://"+bucket+"/"+*value.Key)
	}
	return ret, nil
}

func (finder *S3Finder) Download(src, dst string) (int64, error) {
	stat := prome.NewStat("S3Finder.Download")
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

	file, err := os.Create(dst)
	if err != nil {
		stat.MarkErr()
		return 0, err
	}
	client := s3manager.NewDownloader(finder.Session,
		func(d *s3manager.Downloader) {
			d.Concurrency = 1 //控制下载速度
		})
	timeout := 60
	if finder.Config.Timeout > 0 {
		timeout = finder.Config.Timeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	length, err := client.DownloadWithContext(ctx, file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(path),
		})
	if err != nil {
		stat.MarkErr()
		return 0, err
	}
	zlog.LOG.Info("S3 Finder", zap.String("path", src), zap.Int64("length", length))
	return length, nil
}

func (finder *S3Finder) GetUpdateTime(filepath string) int64 {
	stat := prome.NewStat("S3Finder.GetUpdateTime")
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
	client := s3.New(finder.Session)
	object, err := client.HeadObject(
		&s3.HeadObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(path),
		})
	if err != nil {
		stat.MarkErr()
		return -1
	}
	return object.LastModified.Unix()
}

func (finder *S3Finder) GetETag(filepath string) string {
	stat := prome.NewStat("S3Finder.GetETag")
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
	client := s3.New(finder.Session)
	object, err := client.HeadObject(
		&s3.HeadObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(path),
		})
	if err != nil {
		stat.MarkErr()
		return ""
	}
	etag := *object.ETag
	return etag[1 : len(etag)-1]
}
