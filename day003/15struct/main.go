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
	addr address // 嵌套背的结构体
}

func main() {
	var stu1 = student{
		name: "李三",
		age:  18,
		addr: address{
			province: "北京",
			city:     "厦门",
		},
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.addr)
	fmt.Println(stu1.addr.province)
	fmt.Println(stu1.addr.city)
}
