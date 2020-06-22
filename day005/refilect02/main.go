package main

import (
	"fmt"
	"reflect"
)

// 反射TypeOf()
func reflectType(x interface{}) {
	t := reflect.TypeOf(x) // 拿到x的动态类型信息
	fmt.Printf("type:%v\n", t)
	// fmt.Printf("%T\n", x) // 反射 原理就是用反射，代码补全也是用的反射
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	//指针类型的t.Name()是空
}

type cat struct {
	name string
}
type person struct {
	age  uint8
	name string
}

func main() {
	// reflectType(100)
	// reflectType(13.245)
	// reflectType(false)
	// reflectType("沙河又有沙河")
	// reflectType([3]int{1, 2, 3})
	// reflectType(map[string]int{})

	// 测试自定结构体类型
	var c1 = cat{
		name: "花花",
	}
	var c2 = person{
		name: "李三",
		age:  34,
	}

	reflectType(c1)
	reflectType(c2)
	var i int32 = 100
	var f float32 = 12.34

	reflectType(&i)
	reflectType(&f)

	var a = []int{1, 2, 3}
	reflectType(a)
	var b = [3]int{1, 2, 3}
	reflectType(b)
	reflectType(map[string]int{})
}
