package main

import "fmt"

//  结构体
//  创建新的累心要使用type关键字
type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var haojie = student{
		name:   "张三",
		age:    32,
		gender: "男",
		hobby:  []string{"篮球", "足球", "排球"},
	}
	fmt.Println(haojie)
	fmt.Println(haojie.name)
	fmt.Println(haojie.age)
	fmt.Println(haojie.gender)
	fmt.Println(haojie.hobby)

}
