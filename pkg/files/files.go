package files

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/lylex/drm/pkg/utils"
)

// GetWd is used to get the current directory.
// If any error raises, application will exist.
func GetWd() string {
	dir, err := os.Getwd()
	if err != nil {
		utils.ErrExit("error getting current directory: %s\n", err)
	}

	return dir
}

// IsAbsolutePath is used to judge whether a given path is a absolute path.
// The logic is simple, if a path is start with '/', then it is a absolute path.
func IsAbsolutePath(path string) bool {
	return strings.HasPrefix(path, "/")
}

// IsExist is used to assert whether a file or directory in the path exists or not.
func IsExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		utils.ErrExit("error getting file info: %s\n", err)
	}
	return true
}

// IsDir is used to assert whether the path is a directory or not.
func IsDir(path string) bool {
	return getFileInfo(path).IsDir()
}

// Move is used to move a file or directory from a absolute path to another.
// If the file name differs, it can also act as changing filename.
func Move(src, dir string) {
	if err := os.Rename(src, dir); err != nil {
		utils.ErrExit("error moving file or directory: %s\n", err)
	}
}

// Mkdir create a directry.
func Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Name is used to retrieve the basename of the file or directory of the given path.
func Name(path string) string {
	return filepath.Base(path)
}

// Dir retrieves the Directory path of the passed in full path.
func Dir(fullPath string) string {
	return filepath.Dir(fullPath)
}

func getFileInfo(path string) os.FileInfo {
	fileInfo, err := os.Stat(path)
	if err != nil {
		utils.ErrExit("error getting file info: %s\n", err)
	}
	return fileInfo
}
