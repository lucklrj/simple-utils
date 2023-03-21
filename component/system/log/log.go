package log

import (
	"os"
	"path/filepath"

	"github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func Run(logPath string) error {
	binPath, err := os.Executable()
	if err != nil {
		return err
	}

	logPath = filepath.Dir(binPath) + "/" + logPath

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
