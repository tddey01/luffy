package main

import "fmt"

// 函数
func f1() {
	fmt.Println("hello 沙河")
}
func f2(name1, name2 string) {
	fmt.Println(name1)
	fmt.Println(name2)
}

// 可变参数 0个或多个
func f3(names ...string) {
	fmt.Println(names) // []string
}

//  无返回值
func f4() {
	fmt.Println("MJJ真烦")
}
func f5(a, b int) int {
	return a + b
}

// 多返回值, 必须用括号括起来，用英文逗号分隔
func f6(a, b int) (int, int) {
	return a + b, a - b
}

// 命名的返回值
func f7(a, b int) (sum int, sub int) {
	sum, sub = a+b, a-b
	return
}

//匿名函数
var f = func(name string) {
	fmt.Println("Hello", name)
}

//  闭包 函数包含调用包含外层的变量
func closure(key string) func(name string) {
	key = "北京"
	return func(name string) {
		fmt.Println("hello", name)
		fmt.Println(key)
	}
}

func main() {
	// n1, n2 := f7(10, 5)
	// fmt.Println(n1, n2)
	// //  通过变量调用匿名函数
	// f("豪杰")
	// //  声明并直接调用函数
	// func(name string) {
	// 	fmt.Println("Hello", name)
	// }("北京")

	// 闭包
	f := closure("kevin") // 得到闭包函数
	fmt.Printf("%T\n", f)
	f("dk")             // 调用闭包函数
	f2 := closure("上海") // 得到闭包函数
	f2("豪杰")            // 调用闭包函数
}
