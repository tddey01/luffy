package main

import "fmt"

func main() {
	// //  声明
	var a [5]int  // 定义一个长度为5存放int类型的数组
	var b [10]int // 定义一个长度为10存放int类型的数组

	// 初始化 赋值
	a = [5]int{1, 2, 3, 4, 5}
	b = [10]int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(b)

	// // var c =[3]string = [3]string{"北京","上海","深圳"}
	// var c = [3]string{"北京", "上海", "深圳"}
	// fmt.Println(c)
	// // ... 表示让编译器去数下有多少个初始值，然后给变量赋值类型
	// var d = [...]int{1, 23, 3, 4555, 3453, 6436, 1, 4, 5, 6, 7, 8, 2}
	// fmt.Println(d)
	// fmt.Printf("c:%T d:%T\n", c, d)

	// //  根据索引值初始化
	var e [100]int
	e = [100]int{99: 1}
	fmt.Println(e)

	// 数组的基本使用
	fmt.Println(e[99])
	// 遍历数组的方式1
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// 第二种方式
	// for index, value := range a {
	// 	fmt.Println(index, a[value], value)
	// }
}
