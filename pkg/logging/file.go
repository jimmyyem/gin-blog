package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath   = "runtime/logs/"
	LogSaveName   = "log"
	LogFileExt    = "log"
	LogSqlFileExt = "sql.log"
	TimeFormat    = "20060102"
)

// 保存日志的目录
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// 保存日志的文件路径
func getLogFileFullPath(ftype string) string {
	var fileExt string
	switch ftype {
	case "sql":
		fileExt = LogSqlFileExt
	default:
		fileExt = LogFileExt
	}

	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), fileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// 打开文件返回 *os.File
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

// 创建目录（目录不存在时调用）
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
