package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

//日志操作
//获取日志文件路径
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

//获取日志文件完整路径
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

//打开日志文件
func openLogFile(filepath string) *os.File {
	//判断文件路径是否错误
	_, err := os.Stat(filepath)
	//分文件类型进行处理
	switch {
	//文件夹不存在
	case os.IsNotExist(err):
		mKDir()
	//操作是否允许
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}
	//打开日志文件，追加，创建，
	handle, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

//创建文件
func mKDir() {
	//根路径
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
