package finder

import (
	"github.com/uopensail/ulib/commonconfig"
)

type IFinder interface {
	ListDir(dir string) ([]string, error)
	Download(src, dst string) (int64, error)
	GetUpdateTime(filepath string) int64
	GetETag(filepath string) string
}

func GetFinder(cfg *commonconfig.FinderConfig) IFinder {
	switch cfg.Type {
	case "s3":
		return NewS3Finder(cfg)
	case "oss":
		return NewOSSFinder(cfg)
	case "local":
		return NewLocalFinder(cfg)
	default:
		return nil
	}
}
