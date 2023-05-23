package system

import (
	"os"
	"path/filepath"
)

func IsFileExist(path string) bool {
	s, err := os.Stat(path)
	if err == nil {
		if os.IsExist(err) && !s.IsDir() {
			return true
		} else {
			return false
		}
	}
	return false
}

func IsDirExist(path string) bool {
	s, err := os.Stat(path)
	if err == nil {
		if os.IsExist(err) && s.IsDir() {
			return true
		} else {
			return false
		}
	}
	return false
}
func MakeDirs(path string) error {
	return os.MkdirAll(path, 0777)
}

func GetBinDir() (string, error) {
	binPath, err := os.Executable()
	if err != nil {
		return "", err
	} else {
		return filepath.Dir(binPath), nil
	}

}
