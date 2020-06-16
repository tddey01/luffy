package main

import "fmt"

// 给结构体内嵌模拟“继承”
type animal struct {
	name string
}

// 定义一个动物会动的方法
func (a *animal) move() {
	fmt.Printf("%s 会动~\n", a.name)
}

//  定义一个狗的结构体
type dog struct {
	feet int
	*animal
}

//  定义了一个狗的方法 wangwang
func (d *dog) wangwang() {
	fmt.Printf("%s 谁在: 旺旺~\n", d.name)
}
func main() {
	var a = dog{
		feet: 4,
		animal: &animal{
			name: "旺财",
		},
	}
	a.wangwang() //调用狗的方法
	a.move()     //调用动物的方法
}
