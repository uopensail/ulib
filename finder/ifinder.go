package finder

import (
	"github.com/uopensail/ulib/commonconfig"
)

type IFinder interface {
	ListDir(dir string) ([]string, error)
	Download(src, dst string) (int64, error)
	GetUpdateTime(filepath string) int64
}

func GetFinder(cfg *commonconfig.FinderConfig) IFinder {
	switch cfg.Type {
	case "s3":
		return NewS3Finder(cfg)
	case "oss":
		return NewOSSFinder(cfg)
	default:
		return nil
	}
}
