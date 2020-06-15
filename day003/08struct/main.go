package main

import "fmt"

//  结构体是一个值类型
type student struct {
	name string
	age  int8
}

func main() {
	var stu1 = student{
		name: "豪杰",
		age:  18,
	}
	var stu2 = stu1
	stu2.name = "王三"
	fmt.Println(stu1.name)
	fmt.Println(stu2.name)

	stu3 := &stu1 // 将stu1对应的地址赋值给了stu3， stu3的类型是一个*student
	fmt.Println("%T\n", stu3)
	(*stu3).name = "李四"
	fmt.Println(stu1.name, stu2.name, stu3.name)

}
