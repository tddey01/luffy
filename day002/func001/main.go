package main

import "fmt"

// 函数
// 定义一个不需要参数也没有返回值的函数： sayHello
func sayHello() {
	fmt.Println("hello 上海")
}

// 定义一个接收string类型的name参数
func sayhello2(name string) {
	fmt.Println("sayhello2", name)
}

// 定义接收多个参数的函数并且有一个返回值
// func intSum(a, b int) int {
// 	ret := a + b
// 	return ret
// }
func intSum(a int, b int) (ret int) {
	ret = a + b
	return
}

// 函数接收可变参数,在参数后面加... 表示可变参数
// 可变参数在函数中是切片类型
func intSum3(a ...int) int {
	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	// fmt.Printf("%T\n", a)
	return ret
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
func intSum4(a int, b ...int) int {
	ret := a
	for _, arg := range b {
		ret = ret + arg
	}
	// fmt.Printf("%T\n", a)
	return ret
}

//  go语言中函数类型简写
func intSum5(a, b int) (ret int) {
	ret = a + b
	return
}

// 定义有多个返回值的函数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	//函数调用
	// sayHello()
	// name := "上海"
	// sayhello2(name)
	// sayhello2("深圳")
	// sun := intSum(3, 5)
	// fmt.Println(sun)
	// sun := intSum3(10, 20, 40)
	// fmt.Println(sun)

	// sun := intSum4(10, 20, 40)
	// fmt.Println(sun)
	sum, sub := calc(100, 200)
	fmt.Println(sum, sub)7
}
