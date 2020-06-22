package main

import "fmt"

// 接口值
func main() {
	var x interface{} //  <Type,Value>
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1

	x = a // <int64 , 100>
	fmt.Printf("%T %v\n", x, x)
	x = b // <int32 , 10>
	fmt.Printf("%T %v\n", x, x)
	x = c // <int8 , 1>
	fmt.Printf("%T %v\n", x, x)
	x = 12.34 //<float64, 12.34>
	fmt.Printf("%T %v\n", x, x)
	// x = false //<bool, false>
	// fmt.Printf("%T %v\n", x, x)

	// 类型断言(猜)
	// 如果猜对了 ok=true  value=对应的类型值
	// 如果猜错了， ok=false value=对应类型的零值
	// value, ok := x.([3]int) //数组
	value, ok := x.(interface{})
	fmt.Printf("ok:%t value:%#v value:%T\n", ok, value, value)
	// if ok {
	// 	fmt.Printf("x存到是一个init类型， 值是%v\n", value)
	// } else {
	// 	fmt.Println("x存的不是init类型")
	// }

}
