package utils

// logger.go 日志模块
import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	// 创建日志文件夹
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		err := os.Mkdir("./log", 0755)
		if err != nil {
			logrus.Fatalf("Failed to create log directory: %v", err)
		}
	}

	// 初始化全局logger实例
	Log = logrus.New()

	// 设置日志格式为JSON
	Log.SetFormatter(&logrus.JSONFormatter{})

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./log/info.log",
		logrus.WarnLevel:  "./log/warn.log",
		logrus.ErrorLevel: "./log/error.log",
	}
	Log.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	// 确保日志文件存在
	ensureLogFileExists("./log/info.log")
	ensureLogFileExists("./log/warn.log")
	ensureLogFileExists("./log/error.log")

	// 设置日志级别
	Log.SetLevel(logrus.InfoLevel)

	// 设置输出目标为对应文件
	Log.SetOutput(os.Stdout)
}

// 确保日志文件存在
func ensureLogFileExists(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Fatalf("Failed to create log file: %s, error: %v", filePath, err)
		}
		err = file.Close()
		if err != nil {
			logrus.Fatalf("Failed to close log file: %s, error: %v", filePath, err)
		}
	}
}
