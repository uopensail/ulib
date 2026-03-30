// Package targz provides methods to create and extract tar.gz archives.
//
// Usage:
//
//	targz.Compress("path/to/the/directory/to/compress", "my_archive.tar.gz")
//	targz.Extract("my_archive.tar.gz", "directory/to/extract/to")
//
// Compress creates an archive at my_archive.tar.gz containing the last
// directory component of inputFilePath. Extract restores that directory
// under the given output path, creating any missing parent directories.
package targz

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
)

// Compress creates a .tar.gz archive from the directory at inputFilePath and
// writes it to outputFilePath. Only the last directory component of
// inputFilePath is stored in the archive, not the full path. Parent
// directories of outputFilePath are created if they do not exist. On failure
// any newly created output directories are removed to leave the filesystem
// clean.
func Compress(inputFilePath, outputFilePath string) (err error) {
	inputFilePath = stripTrailingSlashes(inputFilePath)
	inputFilePath, outputFilePath, err = makeAbsolute(inputFilePath, outputFilePath)
	if err != nil {
		return err
	}

	undoDir, err := mkdirAll(filepath.Dir(outputFilePath), 0755)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			undoDir()
		}
	}()

	return compress(inputFilePath, outputFilePath, filepath.Dir(inputFilePath))
}

// Extract unpacks the .tar.gz archive at inputFilePath into the directory
// outputFilePath. Missing parent directories are created automatically. On
// failure any newly created output directories are removed.
func Extract(inputFilePath, outputFilePath string) (err error) {
	outputFilePath = stripTrailingSlashes(outputFilePath)
	inputFilePath, outputFilePath, err = makeAbsolute(inputFilePath, outputFilePath)
	if err != nil {
		return err
	}

	undoDir, err := mkdirAll(outputFilePath, 0755)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			undoDir()
		}
	}()

	return extract(inputFilePath, outputFilePath)
}

// mkdirAll creates dirPath and all necessary parents with the given
// permissions. It returns a cleanup function that removes the first
// directory component that was newly created, so callers can undo the
// creation on error. If the directory already exists the returned function
// is a no-op.
func mkdirAll(dirPath string, perm os.FileMode) (func(), error) {
	var undoDir string

	for p := dirPath; ; p = path.Dir(p) {
		finfo, err := os.Stat(p)
		if err == nil {
			if finfo.IsDir() {
				break
			}
			// p exists but is not a directory; check via Lstat in case it
			// is a symlink pointing to a directory.
			finfo, err = os.Lstat(p)
			if err != nil {
				return nil, err
			}
			if finfo.IsDir() {
				break
			}
			return nil, &os.PathError{Op: "mkdirAll", Path: p, Err: syscall.ENOTDIR}
		}
		if os.IsNotExist(err) {
			undoDir = p
		} else {
			return nil, err
		}
	}

	if undoDir == "" {
		return func() {}, nil
	}

	if err := os.MkdirAll(dirPath, perm); err != nil {
		return nil, err
	}

	return func() { os.RemoveAll(undoDir) }, nil
}

// stripTrailingSlashes removes a single trailing slash from path if present.
func stripTrailingSlashes(p string) string {
	return strings.TrimRight(p, "/")
}

// makeAbsolute converts both paths to absolute form using filepath.Abs.
func makeAbsolute(inputFilePath, outputFilePath string) (string, string, error) {
	absIn, err := filepath.Abs(inputFilePath)
	if err != nil {
		return inputFilePath, outputFilePath, err
	}
	absOut, err := filepath.Abs(outputFilePath)
	return absIn, absOut, err
}

// compress creates the .tar.gz archive. subPath is stripped from each file's
// path so the archive contains paths relative to the directory being
// compressed. The output file is removed if any error occurs after creation.
func compress(inPath, outFilePath, subPath string) (err error) {
	entries, err := os.ReadDir(inPath)
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return errors.New("targz: input directory is empty")
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			os.Remove(outFilePath)
		}
	}()

	// Layer writers: file → gzip → tar
	gzipWriter := gzip.NewWriter(outFile)
	tarWriter := tar.NewWriter(gzipWriter)

	if err = writeDirectory(inPath, tarWriter, subPath); err != nil {
		return err
	}
	// Close in reverse order; errors from Close matter for gzip flush.
	if err = tarWriter.Close(); err != nil {
		return err
	}
	if err = gzipWriter.Close(); err != nil {
		return err
	}
	return outFile.Close()
}

// writeDirectory recursively walks directory and writes every file it
// contains to tarWriter. subPath is used to compute archive-relative names.
func writeDirectory(directory string, tarWriter *tar.Writer, subPath string) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		currentPath := filepath.Join(directory, entry.Name())
		if entry.IsDir() {
			if err := writeDirectory(currentPath, tarWriter, subPath); err != nil {
				return err
			}
			continue
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if err := writeTarGz(currentPath, tarWriter, info, subPath); err != nil {
			return err
		}
	}
	return nil
}

// writeTarGz writes a single file to tarWriter. The header name is the file's
// absolute path with subPath stripped from the front, producing an
// archive-relative path. Symlinks are resolved to determine the link target
// stored in the header.
func writeTarGz(filePath string, tarWriter *tar.Writer, fileInfo os.FileInfo, subPath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	evaledPath, err := filepath.EvalSymlinks(filePath)
	if err != nil {
		return err
	}
	evaledSub, err := filepath.EvalSymlinks(subPath)
	if err != nil {
		return err
	}

	// link is non-empty only when filePath itself is a symlink.
	link := ""
	if evaledPath != filePath {
		link = evaledPath
	}

	header, err := tar.FileInfoHeader(fileInfo, link)
	if err != nil {
		return err
	}

	// Ensure evaledPath starts with evaledSub before slicing to avoid
	// a panic or garbled name on unexpected filesystem layouts.
	if !strings.HasPrefix(evaledPath, evaledSub) {
		return fmt.Errorf("targz: file path %q is not under sub path %q", evaledPath, evaledSub)
	}
	header.Name = evaledPath[len(evaledSub):]

	if err = tarWriter.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, file)
	return err
}

// extract decompresses the .tar.gz archive at filePath into directory.
// Each entry's parent directories are created as needed. Files are written
// with a buffered writer for efficiency and closed explicitly so that any
// flush error is returned to the caller rather than silently lost.
func extract(filePath, directory string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// gzip.NewReader performs its own buffering; an extra bufio layer is
	// unnecessary and only adds overhead.
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Guard against path traversal attacks ("../../../etc/passwd" style).
		cleanName := filepath.Clean(header.Name)
		if strings.HasPrefix(cleanName, "..") {
			return fmt.Errorf("targz: illegal file path in archive: %q", header.Name)
		}

		targetDir := filepath.Join(directory, filepath.Dir(cleanName))
		targetPath := filepath.Join(targetDir, filepath.Base(cleanName))

		if err = os.MkdirAll(targetDir, 0755); err != nil {
			return err
		}

		if err = extractFile(tarReader, targetPath, header.FileInfo().Mode()); err != nil {
			return err
		}
	}

	return nil
}

// extractFile writes the current tar entry from r into a new file at path
// with the given permission bits. The file is closed explicitly so that
// bufio.Writer.Flush errors are propagated correctly.
func extractFile(r io.Reader, path string, mode os.FileMode) (err error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(f, r)
	return err
}
