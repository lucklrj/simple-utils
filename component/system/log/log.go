package log

import (
	"os"
	"path/filepath"

	"github.com/lestrrat-go/file-rotatelogs"
	systemHelper "github.com/lucklrj/simple-utils/helper/system"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func Run(logPath string) error {
	binPath, err := os.Executable()
	if err != nil {
		return err
	}

	logPath = filepath.Dir(binPath) + "/" + logPath
	if systemHelper.IsDirExist(logPath) == false {
		err := systemHelper.MakeDirs(logPath)
		if err != nil {
			return err
		}
	}
	writer, err := rotatelogs.New(
		logPath+"/%Y-%m-%d.log",
		rotatelogs.WithRotationTime(1),
	)
	if err != nil {
		return err
	}
	Logger = log.New()
	Logger.SetOutput(writer)
	Logger.SetFormatter(&log.JSONFormatter{})

	return nil
}
