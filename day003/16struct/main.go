package main

import "fmt"

// 给结构体内嵌模拟“继承”
type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Println("%s会动~")
}

type dog struct {
	feet int
}

func main() {

}
