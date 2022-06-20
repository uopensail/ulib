package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
)

const filechunk = 8192

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
	if len(dir) == 0 {
		return nil, fmt.Errorf("dir: %s error", dir)
	}

	if dir[len(dir)-1] == '/' {
		dir = dir[:len(dir)-1]
	}
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(files))
	for _, f := range files {
		ret = append(ret, dir+"/"+f.Name())
	}
	return ret, nil
}

func GetMD5(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer file.Close()

	// calculate the file size
	info, _ := file.Stat()
	filesize := info.Size()

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))
	hash := md5.New()
	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)
		file.Read(buf)
		io.WriteString(hash, string(buf))
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}
