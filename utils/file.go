package utils

import (
	"io"
	"io/ioutil"
	"os"
)

func FilePathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetFileModifyTime(fileName string) int64 {
	info, err := os.Stat(fileName)
	if err != nil {
		return -1
	}
	return info.ModTime().Unix()
}

func GetFileSize(fileName string) int64 {
	info, err := os.Stat(fileName)
	if err != nil {
		return -1
	}
	return info.Size()
}

func CopyFile(src, dst string) (int64, error) {
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func ListDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(files))
	for _, f := range files {
		ret = append(ret, f.Name())
	}
	return ret, nil
}
