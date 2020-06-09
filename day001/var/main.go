package main

import (
	"fmt"
)

//全部变量声明使用   var
var alix = "abs"

func foo() (string, int) {
	return "alex", 9000
}

func main() {
	// var name string
	// var age int
	// // 声明的变量必须要使用
	// fmt.Println(name)
	// fmt.Println(age)

	// // 批量声明变量
	// var (
	// 	a string
	// 	b int
	// 	c bool
	// 	d string
	// )

	// a = "沙河"
	// b = 1
	// c = true
	// d = "老李"
	// fmt.Println(a, b, c, d)
	// //什么变量 并且初始化值
	// var x string = "老男孩"
	// fmt.Println(x)
	// fmt.Printf("%s教育\n", x)
	// // 类型推导(编译器根基初始值的类型，给指定变量)
	// var y = 200
	// var z = true
	// fmt.Println(y)
	// fmt.Println(z)

	// // 简短变量声明：（只能在函数内部使用）
	// nazha := "嘿嘿"
	// fmt.Println(nazha)
	// 调用foo函数
	// _ （匿名变量）用于接收不需要的变量值
	aa, _ := foo()
	fmt.Println(aa)
}
