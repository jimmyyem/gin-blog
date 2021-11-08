package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Level int

var (
	out1, out2 *os.File
	Logger     *logrus.Logger
	SqlLogger  *logrus.Logger
)

func init() {
	// 普通日志
	filePath := getLogFileFullPath("")
	out1 = openLogFile(filePath)
	Logger = logrus.New()
	Logger.SetOutput(out1)

	// SQL日志
	filePath = getLogFileFullPath("sql")
	out2 = openLogFile(filePath)
	SqlLogger = logrus.New()
	SqlLogger.SetOutput(out2)

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)
	SqlLogger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	SqlLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func Debug(v ...interface{}) {
	Logger.Debug(v)
}

func Info(v ...interface{}) {
	Logger.Info(v)
}

func Warn(v ...interface{}) {
	Logger.Warn(v)
}

func Error(v ...interface{}) {
	Logger.Error(v)
}

func Fatal(v ...interface{}) {
	Logger.Fatal(v)
}
