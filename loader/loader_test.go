package loader

import (
	"fmt"
	"github.com/uopensail/ulib/commonconfig"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoader(t *testing.T) {
	cfg := &commonconfig.DownloaderConfig{
		SourcePath: "/tmp/a.txt",
		LocalPath:  "/tmp/b.txt",
		Interval:   20,
	}

	cfg.Type = "local"
	createFunc := func(path string, p interface{}) ITable {
		file, _ := os.Open(path)
		data, _ := ioutil.ReadAll(file)
		fmt.Println(string(data))
		return file
	}

	releaseFunc := func(table ITable, param interface{}) {
		table.(*os.File).Close()
	}

	Register("a", cfg, createFunc, releaseFunc, cfg.LocalPath, nil)
	select {}
}
