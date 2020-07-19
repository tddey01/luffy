package main

// 用来专门定义与数据库结构体

// Book 书结构体
type Book struct {
	ID    int64   `db:"id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
	Publisher
}

// Publisher 作者 地区
type Publisher struct {
	ID       int64  `db:"id"`
	Province string `db:"province"`
	City     string `db:"city"`
}
