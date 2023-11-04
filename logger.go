package tortorCoin

// logger.go

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	// 初始化全局logger实例
	Log = logrus.New()

	// 设置日志输出格式为JSON
	Log.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别
	Log.SetLevel(logrus.InfoLevel)

	// 设置输出目标为标准输出，也可以设置为文件等
	Log.SetOutput(os.Stdout)
}
