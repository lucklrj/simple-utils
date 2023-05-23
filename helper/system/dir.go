package system

import (
	"os"
	"path/filepath"
)

func IsDirExist(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func IsFileExist(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()

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
