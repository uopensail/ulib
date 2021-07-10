package finder

import (
	"fmt"
	"github.com/uopensail/ulib/commonconfig"
	"testing"
)

func Test_s3(t *testing.T) {
	dw := &commonconfig.FinderConfig{
		Timeout:   0,
		Endpoint:  "",
		Region:    "XXXXXXXXXXXXXX",
		AccessKey: "XXXXXXXXXXXXXX",
		SecretKey: "XXXXXXXXXXXXXX",
	}

	aws := NewS3Finder(dw)
	path := "XXXXXXXXXXXXXX"
	filepath, _ := aws.getPath(path)
	bucket, _ := aws.getBucket(path)

	fmt.Printf("%s %s\n", filepath, bucket)
	fmt.Print(aws.GetUpdateTime(path))
	//aws.Download(path, "/tmp/s3.db")
	files, err := aws.ListDir("XXXXXXXXXXXXXX")
	fmt.Printf("%v %v\n", files, err)
}
