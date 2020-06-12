package main

import "fmt"

// 匿名函数和闭包

// 定义一个函数他的返回值是一个函数
// 把函数作为返回值
func a(name string) func() {
	// name := "北京"
	return func() {
		fmt.Println("上海", name)
	}
}

func main() {
	// func() {
	// 	fmt.Println("匿名函数")
	// }()

	// 闭包 = 函数 + 外层变量的引用
	r := a("百度") // r此时就是一个闭包
	r()      // 相当于执行了a函数内部的匿名函数
}
