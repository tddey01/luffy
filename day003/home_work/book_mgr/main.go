package main

import (
	"fmt"
	"os"
)

// 使用函数实现一个简单的图书管理系统。
// 每本书有书名、作者、价格、上架信息，
// 用户可以在控制台添加书籍、修改书籍信息、打印所有的书籍列表。

// Book is a struct
type Book struct {
	title   string
	author  string
	price   float32
	publish bool
}

// 主菜单
func showMem() {
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 修改书籍信息")
	fmt.Println("3. 展示所有书籍")
	fmt.Println("4. 退出")
	fmt.Println()
}


var (
	// AllBooks is a slice of Book pointer
	AllBooks []*Book
)

func main() {
	for {
		showMem()
		var input int
		fmt.Scanf("%d\n", &input) //从终端获取输入的数字
		switch input {
		case 1:
			AddBook()
			fmt.Println("添加书籍")
			// AddBook()
		case 2:
			fmt.Println("修改书籍")
			ModifyBook()
		case 3:
			fmt.Println("显示所有书籍")
			ShowAllBook()
		case 4:
			fmt.Println("退出")
			os.Exit(0)

		}
	}
}
