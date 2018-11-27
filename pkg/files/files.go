package files

import (
	"os"
	"strings"

	"github.com/lylex/drm/pkg/utils"
)

// GetWd is used to get the current directory.
// If any error raises, application will exist.
func GetWd() string {
	dir, err := os.Getwd()
	if err != nil {
		utils.PrintErr("error getting current directory: %s", err)
	}

	return dir
}

// IsAbsolutePath is used to judge whether a given path is a absolute path.
// The logic is simple, if a path is start with '/', then it is a absolute path.
func IsAbsolutePath(path string) bool {
	return strings.HasPrefix(path, "/")
}
