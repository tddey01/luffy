package main

import (
	"fmt"
)

// const pi = 3.14

// // 批量声明常量
// const (
// 	a = 100
// 	b = 1000
// 	c
// 	d
// 	// const声明如果不写，默认就和上一行一样
// )

// // iota 每局
// const (
// 	aa = iota
// 	bb = iota
// 	cc
// 	dd
// )

// // 面试题
// const (
// 	n1 = iota
// 	n2
// 	_
// 	n4
// )

// const (
// 	n1 = iota //0
// 	n2 = 100  //100
// 	n3 = iota //1
// 	n4        //3
// )

// const n5 = iota //0

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 // 1 2
	c, d                      // 2 3
	e, f                      // 3 4
)

func main() {
	// // 常量  常量是不允许修改值的
	// fmt.Println(pi)

	// fmt.Println(a, b, c, d)

	// fmt.Println(aa, bb, cc, dd)

	// fmt.Println(n1, n2, n4)
	// fmt.Println(n1, n2, n3, n4, n5)
	// fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e, f)
}

// iota
// 0. const声明如果不写，默认就和上一行一样
// 1. 遇到const iota就初始化为零
// 2. const中每新增一行变量声明iota就递增1
