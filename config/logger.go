package config

import (
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New()

func InitLogger(conf *toml.Tree) {
	logPath := conf.GetDefault("path", "./logs/").(string)
	_, err := os.Stat(logPath)
	if err != nil {
		logger.Fatal("log path not exist & it cannot create auto")
	}
	fileName := conf.GetDefault("file", "chs.log").(string)
	logFile, err := os.OpenFile(logPath+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		logger.SetOutput(logFile)
	}
	logger.Info("Init logger success")
}

func Logger() *logrus.Logger {
	return logger
}

func LoggerWithField(key string, value interface{}) *logrus.Logger {
	return logger.WithField(key, value).Logger
}

/**
 * logrus.Fields <=> map[string]interface{}
 */
func LoggerWithFields(fields logrus.Fields) *logrus.Logger {
	return logger.WithFields(fields).Logger
}
