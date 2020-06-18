package main

import "fmt"

// 为什么要用接口inteface

// Cat 空接口
type Cat struct{}

// Say 返回
func (c Cat) Say() string { return "喵喵喵" }

// Dog 空接口
type Dog struct{}

// Say 返回
func (d Dog) Say() string { return "汪汪汪" }

type pig struct{}

// Say 猪
func (p pig) Say() (r string) {
	r = "哼哼哼"
	return
}

// Sayer 接口
type anima interface {
	Say() string
}

func main() {
	// c := Cat{}
	// fmt.Println("猫", c.Say())
	// d := Dog{}
	// fmt.Println("狗", d.Say())
	var anima1List []anima
	c := Cat{} // 造一个猫
	d := Dog{} // 造一个狗
	p := pig{} // 造一个猪
	anima1List = append(anima1List, c, d, p)
	fmt.Println(anima1List)

	//  代码冗余
	for _, item := range anima1List {
		ret := item.Say()
		fmt.Println(ret)
	}
}
