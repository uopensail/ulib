package finder

import (
	"fmt"
	"github.com/uopensail/ulib/commonconfig"
	"testing"
)

func Test_OSS(t *testing.T) {
	dw := &commonconfig.FinderConfig{
		Timeout:   0,
		Endpoint:  "XXXXXXXXXXXXXXX",
		Region:    "",
		AccessKey: "XXXXXXXXXXXXXXX",
		SecretKey: "XXXXXXXXXXXXXXX",
	}

	aws := NewOSSFinder(dw)
	path := "XXXXXXXXXXXXXXX"
	filepath, _ := aws.getPath(path)
	bucket, _ := aws.getBucket(path)

	fmt.Printf("%s %s\n", filepath, bucket)
	fmt.Print(aws.GetUpdateTime(path))
	len, err := aws.Download(path, "/tmp/a.db")
	fmt.Printf("%v %v\n", len, err)
	files, err := aws.ListDir("XXXXXXXXXXXXXXX")
	fmt.Printf("%v %v", files, err)
}
