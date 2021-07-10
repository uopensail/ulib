package utils

import "os"

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
