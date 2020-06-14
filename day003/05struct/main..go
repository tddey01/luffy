package main

import "fmt"

// 自定义类型
type NewInt int

//  类型别名  只从在代码编写过程中，代码编译之后根本不存在haojie
// 提高代码可读性
type haojie = int

// byte uint8

func main() {
	var a NewInt
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b haojie
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
