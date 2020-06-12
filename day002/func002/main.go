package main

import "fmt"

// 函数进阶之变量作用域

// 定义全局变量num
var num = 10

// 定义函数
func testGlobal() {
	num := 100 // 当前函数中没有num这个变量 会去全局找num这个变量 函数中相同名的变量优先级高于全局变量
	name := "上海"
	// 可以在函数中访问全局变量
	//  现在自己函数中查找， 找到了就用自己的函数中的
	// 函数中找不到变量就往外层找全局变量
	fmt.Println("全局变量", num)
	fmt.Println("name", name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {

	// 	testGlobal()
	// 	// 外层不能访问到内层函数的内部变量（局部变量）
	// 	// fmt.Println(name)
	// 	//  变量i此时只在for循环的语句块中生效
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	// fmt.Println(i) // 外层访问不到for语句块中的变量

	// // 函数可以作为变量
	// abc := testGlobal
	// fmt.Printf("%T\n", abc)
	// abc()

	ret := calc(100, 200, add)
	fmt.Println(ret)
	ret1 := calc(100, 200, sub)
	fmt.Println(ret1)
}
