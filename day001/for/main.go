package main

import (
	"fmt"
)

// for循环
func main() {
	age := 18
	for age > 0 { // 相当于别的语言中的while循环
		fmt.Println(age)
		age = age - 1
		age--
	}

	// 无限循环
	// for {
	// 	循环体语句
	// }
	// for循环可以通过break、goto、return、panic语句强制退出循环。

	// for range(键值循环)
	// 	Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：

	// 数组、切片、字符串返回索引和值。
	// map返回键和值。
	// 通道（channel）只返回通道内的值。

	// switch case
	// 使用switch语句可方便地对大量的值进行条件判断。
	switchDemo1()
	testSwitch3()
}
func switchDemo1() {
	finger := 7
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
