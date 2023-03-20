package log

import (
	"github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func Run(logPath string) error {

	writer, err := rotatelogs.New(
		logPath+".%Y-%m-%d",
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
