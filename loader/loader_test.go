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
	foo := func(p interface{}) ITable {
		file, _ := os.Open(p.(string))
		data, _ := ioutil.ReadAll(file)
		fmt.Println(string(data))
		return nil
	}
	Register("a", cfg, foo, cfg.LocalPath)
	Register("b", cfg, foo, cfg.LocalPath)
	Register("c", cfg, foo, cfg.LocalPath)
	Register("d", cfg, foo, cfg.LocalPath)
	Register("e", cfg, foo, cfg.LocalPath)
	select {}
}
