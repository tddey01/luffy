package main

import "fmt"

// NewBook is the func of create new Book instance.
func NewBook(title, author string, price float32, pubilsh bool) *Book {
	return &Book{
		title:   title,
		author:  author,
		price:   price,
		publish: pubilsh,
	}
}

// inputBook 是一个从终端获取输入 返回Book指针的一个函数
func inputBook() *Book {
	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Println("")
	fmt.Println("请输入书名")
	fmt.Scanf("%s\n", &title)
	fmt.Println("请输入作者")
	fmt.Scanf("%s\n", &author)
	fmt.Println("请输入价格")
	fmt.Scanf("%f\n", &price)
	fmt.Println("请输入是和否上架[true|false]")
	fmt.Scanf("%t\n", &publish)
	newBook := NewBook(title, author, price, publish)
	return newBook
}


// AddBook 是添加书籍的方法
func AddBook() {
	newBook := inputBook()
	for _, v := range AllBooks {
		if v.title == newBook.title {
			fmt.Printf("书名为%s的书已经存在！\n", newBook.title)
			return
		}
	}
	AllBooks = append(AllBooks, newBook)
}

// ModifyBook is the func modify the book info.
func ModifyBook() {
	newBook := inputBook()

	for index, v := range AllBooks {
		if v.title == newBook.title {
			AllBooks[index] = newBook
			fmt.Printf("书名为%s的书籍信息更新成功！\n", newBook.title)
			return
		}
	}
	fmt.Printf("根据书名：%s无法查找到书籍信息！\n", newBook.title)
}

// ShowAllBook  is the func of list all books. 
func ShowAllBook() {
	for _, v := range AllBooks {
		fmt.Printf("书名:《%s》 作者:%s 价格:%.2f 是否上架:%t\n", v.title, v.author, v.price, v.publish)
	}
}
