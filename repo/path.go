package repo

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetPath(path string) (string, error) {
	if path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = cwd
	}

	expandedPath, err := expandPath(path)
	if err != nil {
		return "", err
	}

	_, err = os.Stat(expandedPath)
	if os.IsNotExist(err) {
		return "", err
	}

	return expandedPath, nil
}

// expandPath expands the ~ symbol to the home directory of the current user
func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		path = filepath.Join(usr.HomeDir, path[1:])
	}
	return path, nil
}
