package helper

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

func NewLog() *logrus.Logger {
	// 日志文件
	fileName := path.Join(GetAppConf().Log.Path, GetAppConf().Log.Name)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return logger
}
