package main

import "fmt"

//  结构体的嵌套
type address struct {
	province string
	city     string
}

type student struct {
	name string
	age  int
	// addr address // 嵌套背的结构体
	address // 嵌套背的结构体
}

func main() {
	var stu1 = student{
		name: "李三",
		age:  18,
		// addr: address{ //匿名结构体默认是类名作为字段结构体
		address: address{ //匿名结构体默认是类名作为字段结构体
			province: "北京",
			city:     "厦门",
		},
	}
	fmt.Println(stu1.name)
	// fmt.Println(stu1.addr)
	// fmt.Println(stu1.addr.province)
	// fmt.Println(stu1.addr.city)
	fmt.Println(stu1.address)
	fmt.Println(stu1.address.province)
	fmt.Println(stu1.address.city)
	fmt.Println(stu1.province)
}
