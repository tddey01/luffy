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

	// 实例化方法1
	//  struct 是值类型的
	// 如果初始化时没有给属性（字段）设置对应的初始值，那么对应属性就是其默认值
	var wangzhen = student{}
	fmt.Println(wangzhen.name)
	fmt.Println(wangzhen.age)
	fmt.Println(wangzhen.gender)
	fmt.Println(wangzhen.hobby)

	// 实例化方法2 //T:表示类型或者结构体
	var yawei = new(student)
	fmt.Println(yawei)
	// (*yawei).name
	yawei.name = "亚伟"
	yawei.age = 32
	yawei.gender = "女"
	fmt.Println(yawei.name, yawei.age)

	// 实例化方法3
	var nazha = &student{}
	fmt.Println(nazha)
	nazha.name = "娜扎"
	nazha.age = 34
	fmt.Println(nazha.name, nazha.age)

	// 结构体初始化
	var stu1 = student{
		"上海",
		18,
		"男",
		[]string{"男人", "女人"},
	}
	fmt.Println(stu1.name, stu1.age)

	var stu2 = &student{
		name: "豪杰",
		age:  19,
	}
	fmt.Println(stu2.name, stu2.age, stu2.gender)
}
