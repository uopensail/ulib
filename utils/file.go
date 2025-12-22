package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"path/filepath"
)

const filechunk = 8192 // 8KB chunks for file reading

// FilePathExists checks if a file or directory exists at the given path
//
// @param path: The file or directory path to check
// @return: true if path exists, false otherwise
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

// GetFileModifyTime retrieves the last modification time of a file
//
// @param fileName: Path to the file
// @return: Unix timestamp of last modification, -1 if error occurs
func GetFileModifyTime(fileName string) int64 {
	info, err := os.Stat(fileName)
	if err != nil {
		return -1
	}
	return info.ModTime().Unix()
}

// GetFileSize retrieves the size of a file in bytes
//
// @param fileName: Path to the file
// @return: File size in bytes, -1 if error occurs
func GetFileSize(fileName string) int64 {
	info, err := os.Stat(fileName)
	if err != nil {
		return -1
	}
	return info.Size()
}

// CopyFile copies a file from source to destination
//
// @param src: Source file path
// @param dst: Destination file path
// @return: Number of bytes copied and error if operation fails
func CopyFile(src, dst string) (int64, error) {
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	// Create destination directory if it doesn't exist
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return 0, err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// ListDir lists all files and directories in the specified directory
//
// @param dir: Directory path to list
// @return: Slice of full paths to directory entries and error if operation fails
func ListDir(dir string) ([]string, error) {
	if len(dir) == 0 {
		return nil, fmt.Errorf("directory path cannot be empty")
	}

	// Check if directory exists
	if !FilePathExists(dir) {
		return nil, fmt.Errorf("directory does not exist: %s", dir)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0, len(files))
	for _, f := range files {
		ret = append(ret, path.Join(dir, f.Name()))
	}
	return ret, nil
}

// GetMD5 calculates the MD5 hash of a file
//
// @param filepath: Path to the file
// @return: MD5 hash as hexadecimal string, empty string if error occurs
func GetMD5(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer file.Close()

	// Calculate the file size
	info, err := file.Stat()
	if err != nil {
		return ""
	}
	filesize := info.Size()

	// Handle empty files
	if filesize == 0 {
		hash := md5.New()
		return fmt.Sprintf("%x", hash.Sum(nil))
	}

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))
	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		// Read chunk from file
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return ""
		}
		if n == 0 {
			break
		}

		// Write chunk to hash
		if _, err := hash.Write(buf[:n]); err != nil {
			return ""
		}
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

// IsDirectory checks if the given path is a directory
//
// @param path: Path to check
// @return: true if path is a directory, false otherwise
func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// CreateDir creates a directory and all necessary parent directories
//
// @param dir: Directory path to create
// @return: error if operation fails
func CreateDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// RemoveFile removes a file from the filesystem
//
// @param filepath: Path to the file to remove
// @return: error if operation fails
func RemoveFile(filepath string) error {
	return os.Remove(filepath)
}

// RemoveDir removes a directory and all its contents recursively
//
// @param dir: Directory path to remove
// @return: error if operation fails
func RemoveDir(dir string) error {
	return os.RemoveAll(dir)
}

// GetFileExtension gets the extension of a file
//
// @param filename: Name or path of the file
// @return: File extension including the dot (e.g., ".txt")
func GetFileExtension(filename string) string {
	return filepath.Ext(filename)
}

// GetFileNameWithoutExtension gets the filename without extension
//
// @param filename: Path to the file
// @return: Filename without extension
func GetFileNameWithoutExtension(filename string) string {
	basename := filepath.Base(filename)
	ext := filepath.Ext(basename)
	return basename[:len(basename)-len(ext)]
}

// MoveFile moves a file from source to destination
//
// @param src: Source file path
// @param dst: Destination file path
// @return: error if operation fails
func MoveFile(src, dst string) error {
	// Create destination directory if it doesn't exist
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	return os.Rename(src, dst)
}
