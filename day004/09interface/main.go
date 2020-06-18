package main

import (
	"fmt"
)

// 空接口
func showType(a interface{}) {
	fmt.Printf("type:%T\n", a)
}

func main() {
	// 	var x interface{}
	// 	x = 100
	// 	fmt.Println(x)
	// 	x = "上海"
	// 	fmt.Println(x)
	// 	x = false
	// 	fmt.Println(x)
	// 	x = struct{ name string }{name: "芳芳"}
	// 	fmt.Println(x)
	// 	showType(x)
	// 	showType(10.0)
	// 	showType(time.Second)
	// 	showType(time.Now)

	// 定义一个值为空接口map
	var stuInfo = make(map[string]interface{}, 100)
	stuInfo["豪杰"] = 100
	stuInfo["韩星"] = true
	stuInfo["战三"] = 9.99
	stuInfo["李三"] = "呵呵"
	fmt.Println(stuInfo)
}
