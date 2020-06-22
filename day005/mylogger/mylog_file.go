package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志代码

// FIleLogger 日志结构体
type FIleLogger struct {
	Level    Level // 日志级别
	fileName string
	filepath string
	file     *os.File
	errFile  *os.File
	maxSize  int64
}

// NewFileLogger 文件日志结构体构造函数
func NewFileLogger(levelStr, fileName, filepath string) *FIleLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &FIleLogger{
		Level:    logLevel,
		fileName: fileName,
		filepath: filepath,
		maxSize:  10 * 1024 * 1024,
	}
	fl.initFile() // 根据上面的文件路径和文件名打开日志文件， 把文件句柄赋值给结构体字段
	return fl
}

// 将指定的日志文件打开 赋值给结构体
func (f *FIleLogger) initFile() {
	logName := path.Join(f.filepath, f.fileName)
	// 打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志%s文件失败 %v", logName, err))
	}
	f.file = fileObj
	// 打开错误日志的文件

	errLogName := fmt.Sprintf("%s.err.log", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志%s文件失败 %v", errLogName, err))
	}
	f.errFile = errFileObj
}

//  检查是否要拆分
func (f *FIleLogger) checkSplit(file *os.File) bool {
	//  检查当前日志文件大小是否超过了maxSize
	FileInfo, _ := file.Stat()
	fileSize := FileInfo.Size()
	return fileSize >= f.maxSize // 当传进来日志文件大小超过maxSize 就返回true
}

// 封装一个切分日志文件方法
func (f *FIleLogger) splitLoggerFile(file *os.File) *os.File {

	//切分文件
	fileName := file.Name()
	backupName := fmt.Sprintf("%s_%v.log", fileName, time.Now().Unix())
	// 把原来的文件关闭
	f.file.Close()
	// 备份原来的文件
	os.Rename(fileName, backupName)
	//  新建一个文件
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志%s文件失败 %v", fileName, err))
	}
	return fileObj
}

// 将公用记录日志等功能封装成一个单独方法
func (f *FIleLogger) log(level Level, format string, args ...interface{}) {
	if f.Level > level {
		return
	}
	// f.file.Write(b)
	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志

	//  日志格式 [时间][文件:行号][函数名][日志级别]日志信息
	newtime := time.Now().Format("2006-02-03 15:04:05.000")
	fileName, funcName, line := getCallerInfo(3)
	logLevelSt := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s] [%s:%d] [%s] [%s] %s", newtime, fileName, line, funcName, logLevelSt, msg)
	// 往文件里面写之前检查
	if f.checkSplit(f.file) {
		f.file = f.splitLoggerFile(f.file)
	}
	fmt.Fprintln(f.file, logMsg) //利用fmt包净msg字符串写入f.file文件中
	// fmt.Errorf(format, a)
	// fmt.Sprintf(format, a)
	//  如果是error或者fatal级别日志还要记录到发。errFile
	if f.Level >= FATALLevel {
		if f.checkSplit(f.errFile) {
			f.errFile = f.splitLoggerFile(f.errFile)
		}
		fmt.Fprintln(f.errFile, logMsg)
	}
}

// DEBUG 级别日志
func (f *FIleLogger) DEBUG(format string, args ...interface{}) {
	f.log(DEBUGLevel, format, args...)
}

// INFO 级别日志
func (f *FIleLogger) INFO(format string, args ...interface{}) {
	f.log(INFOLevel, format, args...)
}

// WARNING 级别日志
func (f *FIleLogger) WARNING(format string, args ...interface{}) {
	f.log(WARNINGLevel, format, args...)
}

// ERROR 级别日志
func (f *FIleLogger) ERROR(format string, args ...interface{}) {
	f.log(ERRORLevel, format, args...)
}

// FATAL 级别日志
func (f *FIleLogger) FATAL(format string, args ...interface{}) {
	f.log(FATALLevel, format, args...)
}

// Close 关闭日志文件句柄
func (f *FIleLogger) Close() {
	f.file.Close()
	f.errFile.Close()
}
