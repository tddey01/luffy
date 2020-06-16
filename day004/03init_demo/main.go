package main

import "fmt"

// init() 初始化
var today = "星期天"

const Week = 7

type student struct {
	Name string
}

//  包被导入的时候会自动执行(多用来做初始化的操作)

func init() {
	fmt.Println("包被初始化了! ")
	fmt.Println(Week)
}

func main() {
	fmt.Println("这是main函数")
}
