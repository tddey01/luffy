package main

import (
	"fmt"

	"github.com/tddey01/luffy/day004/15log_demo/mylog"
)

//  写了一个项目想要在代码中记录日志
//  要使用mylog这个包
func main() {
	f1 := mylog.NewFileLogger(mylog.DEBUG, "./", "test.log")

	f1.DEBUG("这是一条DEBUG测试日志")
	fmt.Println("可以申请IPO")
	userId := 10
	// f1.INFO("这是一条INFO日志")
	f1.DEBUG("ID是 %d 的用户一直尝试登陆", userId)
}
