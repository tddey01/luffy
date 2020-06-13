package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包

// 定义一个函数他的返回值是一个函数
// 把函数作为返回值
func a(name string) func() {
	// name := "北京"
	return func() {
		fmt.Println("上海", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		return base
	}
	sub := func(i int) int {
		base -= 1
		return base
	}
	return add, sub
}

func main() {
	// func() {
	// 	fmt.Println("匿名函数")
	// }()

	// 闭包 = 函数 + 外层变量的引用
	// r := a("百度") // r此时就是一个闭包
	// r()          // 相当于执行了a函数内部的匿名函数

	// r := makeSuffixFunc(".txt")
	// ret := r("上海")
	// fmt.Println(ret)

	x, y := calc(100)
	ret1 := x(200) // base = 100 +200
	fmt.Println(ret1)
	ret2 := y(200)
	fmt.Println(ret2)

}
