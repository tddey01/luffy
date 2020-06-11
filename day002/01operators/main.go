package main

import (
	"fmt"
)

// 运算符
func main() {
	// n1 := 19
	// n2 := 3
	// 算数运算符
	// fmt.Println(n1 + n2)
	// fmt.Println(n1 - n2)
	// fmt.Println(n1 * n2)
	// fmt.Println(n1 / n2)
	// fmt.Println(n1 % n2)
	// n2++
	// fmt.Println(n2)
	// n1--
	// fmt.Println(n1)

	//关系运算符
	// fmt.Println(n1 == n2)
	// fmt.Println(n1 != n2)
	// fmt.Println(n1 > n2)
	// fmt.Println(n1 >= n2)
	// fmt.Println(n1 < n2)
	// fmt.Println(n1 <= n2)

	// //逻辑运算符
	// a := true
	// b := false
	// // 两个条件都成立为真
	// fmt.Println(a && b)
	// // 逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。 两个条件有一个为真
	// fmt.Println(a || b)
	// // 逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。  原来取费 就为假
	// fmt.Println(!a) //false
	// fmt.Println(!b) //true
	// 位运算
	// 1101 13 11 3
	fmt.Printf("13的二进制%b\n", 13) //1101
	fmt.Printf("3的二进制%b\n", 3)   //11
	//  & (两个对应的二进制位都为1才为1)
	fmt.Println(13 & 3) // 返回1
	//  | 两两个对应的二进制位有一位为1 就唯1
	fmt.Println(13 | 3)
	//  ^ (两个对应的二进制位不一样则为1)
	fmt.Println(13 ^ 3)
	// <<
	fmt.Println(3 << 10)
	//  >>
	fmt.Println(3 >> 1)

}
