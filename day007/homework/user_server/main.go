package main

import (
	"github.com/tddey01/luffy/day007/homework/mylogger"
)

//  测试mylogger的程序
var logger mylogger.Logger

func main() {
	logger = mylogger.NewFileLogger("DEBUG", "xx.log", "./")
	defer logger.Close()
	for {
		logger.DEBUG("下雨了")
		logger.ERROR("打雷了")
	}

}
