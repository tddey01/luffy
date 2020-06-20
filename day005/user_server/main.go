package main

import (
	"github.com/tddey01/luffy/day005/mylogger"
)

var logger mylogger.Logger

//  一个使用自定义日志库的用户程序
func main() {
	logger = mylogger.NewFileLogger("fatal", "./", "xx.log")
	defer logger.Close()
	// logger := mylogger.NewConsoleFileLogger("debug")
	for {
		sb := "馆大门是个好帮"
		logger.DEBUG("%s是很棒的", sb)
		logger.INFO("INOF  这是一条测试的日志")
		logger.ERROR("Error 这是一条测试的日志 ")
		logger.FATAL("Fatal 这是一条测试的日志 ")
	}

}
