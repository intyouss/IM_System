package utils

import (
	"IM_System/config"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var MysqlLogger logger.Interface

func init() {
	mysqlLogConfig := config.GetMysqlLog()
	var logLevel logger.LogLevel
	switch mysqlLogConfig["loglevel"].(string) {
	case "info":
		logLevel = logger.Info
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	}
	MysqlLogger = logger.New(
		log.New(os.Stdout, mysqlLogConfig["prefix"].(string), log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      mysqlLogConfig["colorful"].(bool),
		},
	)
}
