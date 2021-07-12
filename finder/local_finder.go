package finder

import (
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/utils"
)

type LocalFinder struct {
	Config *commonconfig.FinderConfig
}

func NewLocalFinder(conf *commonconfig.FinderConfig) *LocalFinder {
	finder := &LocalFinder{
		Config: conf,
	}
	return finder
}

func (finder *LocalFinder) ListDir(dir string) ([]string, error) {
	stat := prome.NewStat("LocalFinder.ListDir")
	defer stat.End()
	return utils.ListDir(dir)
}

func (finder *LocalFinder) Download(src, dst string) (int64, error) {
	stat := prome.NewStat("LocalFinder.Download")
	defer stat.End()
	return utils.CopyFile(src, dst)
}

func (finder *LocalFinder) GetUpdateTime(filepath string) int64 {
	stat := prome.NewStat("LocalFinder.GetUpdateTime")
	defer stat.End()
	return utils.GetFileModifyTime(filepath)
}

func (finder *LocalFinder) GetMD5(filepath string) string {
	stat := prome.NewStat("LocalFinder.GetMD5")
	defer stat.End()
	return utils.GetMD5(filepath)
}
