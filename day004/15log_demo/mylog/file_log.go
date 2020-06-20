package mylog

import (
	"fmt"
	"os"
	"time"
)

// FileLogger 往文件中记录日志结构体
type FileLogger struct {
	level       int // 只有大于这个级别的日志才会记录，小于不记录
	logFilePath string
	logFileName string
	logFile     *os.File
}

// NewFileLogger 是一个生成日志结构体实例的构造函数
func NewFileLogger(level int, logFilePath, logFileName string) *FileLogger {
	flObj := &FileLogger{
		level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	flObj.initFileLogger() // 调用下面初始化文件句柄的方法
	return flObj
}

// 专门用来初始化日志文件句柄
func (f *FileLogger) initFileLogger() {
	filepath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file:%s failed ", filepath))
	}
	f.logFile = file // 把日志文件复制给结构体中的logFile这个字段
}

// DEBUG  记录日志
func (f *FileLogger) DEBUG(format string, args ...interface{}) {
	if f.level > DEBUG { // 如果你设置日志级别大于当前级别不用写日志
		return
	}
	fileName, funcName, line := getCallerInfo()
	// 往文件里面写
	//  日志的格式要丰富起来 时间 日志级别  那个文件 哪一行  哪一个函数 日志细信息
	// f.logFile.WriteString(msg)
	//  [2020-04-21 18:58:01] [DEBUG] main.go [14] id为 10 的用户一直在尝试登陆
	// nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s][%d] %s", nowStr, getlevelStr(f.level), fileName, funcName, line, format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile) //换行
}

// INFO  记录日志
func (f *FileLogger) INFO(format string, args ...interface{}) {
	if f.level > INFO{
		return
	}
	fileName,funcName,line := getCallerInfo()
	// 往文件里面写
	// f.logFile.WriteString(msg) // 满足不了我们不需求
	newStr :=time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s][%d] %s", newStr, getlevelStr(f.level), fileName, funcName, line, format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile) //换行
}

// ERROR  记录日志
func (f *FileLogger) ERROR(format string, args ...interface{}) {
	if f.level > ERROR{
		return
	}
	fileName,funcName,line := getCallerInfo()
	// 往文件里面写
	// f.logFile.WriteString(msg) // 满足不了我们不需求
	newStr :=time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s][%d] %s", newStr, getlevelStr(f.level), fileName, funcName, line, format)
	// 往文件里面写
	// f.logFile.WriteString(msg)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile) //换行
}
