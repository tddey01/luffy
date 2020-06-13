package main

import "fmt"

// 切片
func main() {
	// // 修改数组的示例
	// //  数组是值类型
	// a1 := [3]int{1, 2, 3}
	// b1 := a1 // al和b1此时都有一个自己对应的【1 2 3】
	// a1[1] = 20
	// fmt.Println(a1) // [1 20,3]
	// fmt.Println(b1) // [1 2 3]

	// //  切片
	// a2 := []int{1, 2, 3}
	// b2 := a2 // 此时a2和b2的小箭头都指向了内存中那个[1 2 3]切片6
	// a2[1] = 20
	// fmt.Println(a2) //【1 2 3】
	// fmt.Println(b2) // [1 2 3]

	// 切片的初始化
	// 1 声明变量时初始化
	//  切片的三要素
	//     第一个元素在底层数组的位置
	//     大小(len) 指的是当前切片中元素的数量
	//     容量（cap)   指的底层数组能容纳的元素的个数
	var x = []int{1, 2, 3} //【1 2 3】
	fmt.Println(x)
	x2 := x[0:2]
	fmt.Println(x2)

	// 2 由一个数组切片得到
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := a[0:len(a)] // a[:]
	fmt.Println(s1)
	s2 := a[0:4] // 从开始一直切到3
	s3 := a[4:]  // 从索引为4一直切到最后
	fmt.Println(s2)
	fmt.Println(s3)

	//  3 由切片在切片得到
	s4 := s2[:]

	fmt.Println(s4)
	// 4 make()创建切片
	s5 := make([]bool, 2, 10)
	fmt.Println(s5)

}
