package mylogger

import (
	"fmt"
	"os"
	"time"
)

//  网终端打印日志

// ConsoleLogger 是一个终端日志结构体
type ConsoleLogger struct {
	Level Level
}

// NewFileLogger 文件日志结构体构造函数
func NewConsoleFileLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	cl := &ConsoleLogger{
		Level: logLevel,
	}
	return cl
}

// 将公用记录日志等功能封装成一个单独方法
func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if c.Level > level {
		return
	}
	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志

	//  日志格式 [时间][文件:行号][函数名][日志级别]日志信息
	newtime := time.Now().Format("2006-02-03 15:04:05.000")
	fileName, funcName, line := getCallerInfo(3)
	logLevelSt := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s] [%s:%d] [%s] [%s] %s", newtime, fileName, line, funcName, logLevelSt, msg)
	fmt.Fprintln(os.Stdout, logMsg) //利用fmt包净msg字符串写入os.Stdout文件中

}

// DEBUG 级别日志
func (c *ConsoleLogger) DEBUG(format string, args ...interface{}) {
	c.log(DEBUGLevel, format, args...)
}

// INFO 级别日志
func (c *ConsoleLogger) INFO(format string, args ...interface{}) {
	c.log(INFOLevel, format, args...)
}

// WARNING 级别日志
func (c *ConsoleLogger) WARNING(format string, args ...interface{}) {
	c.log(WARNINGLevel, format, args...)
}

// ERROR 级别日志
func (c *ConsoleLogger) ERROR(format string, args ...interface{}) {
	c.log(ERRORLevel, format, args...)
}

// FATAL 级别日志
func (c *ConsoleLogger) FATAL(format string, args ...interface{}) {
	c.log(FATALLevel, format, args...)
}

// Close 终端标准输出不需要关闭
func (c *ConsoleLogger) Close() {}
