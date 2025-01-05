package logger

import (
	"os"
	"path/filepath"

	"github.com/babycare/configs"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *logrus.Logger

func InitLogger(config *configs.LogConfig) error {
	Log = logrus.New()

	// 设置日志级别
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return err
	}
	Log.SetLevel(level)

	// 确保日志目录存在
	logDir := filepath.Dir(config.Filename)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// 配置日志输出
	Log.SetOutput(&lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,    // MB
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,     // days
		Compress:   config.Compress,
	})

	// 设置日志格式
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return nil
} 