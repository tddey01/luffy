package mylogger

import (
	"strings"
)

// 我的日志库文件

// Level 是一个自定义的类型代表日志级别
type Level uint16

// Logger 定义一个logger 接口
type Logger interface {
	DEBUG(format string, args ...interface{})
	INFO(format string, args ...interface{})
	WARNING(format string, args ...interface{})
	ERROR(format string, args ...interface{})
	FATAL(format string, args ...interface{})
	Close()
}

// 定义具体日志级别常量
const (
	DEBUGLevel Level = iota
	INFOLevel
	WARNINGLevel
	ERRORLevel
	FATALLevel
)

// 写一个根据传进来的Level 获取对应的字符串
func getLevelStr(level Level) string {
	switch level {
	case DEBUGLevel:
		return "DEBUG"
	case INFOLevel:
		return "INFO"
	case WARNINGLevel:
		return "WARNING"
	case ERRORLevel:
		return "ERROR"
	case FATALLevel:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

//  根据用户传入的字符类型的日志级别， 解析出对应的Level
func parseLogLevel(levelStr string) Level {
	levelStr = strings.ToLower(levelStr) // 将字符串转成全部小写
	switch levelStr {
	case "debug":
		return DEBUGLevel
	case "info":
		return INFOLevel
	case "warning":
		return WARNINGLevel
	case "error":
		return ERRORLevel
	case "fatal":
		return FATALLevel
	default:
		return DEBUGLevel
	}
}
