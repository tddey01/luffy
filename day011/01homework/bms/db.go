package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 数据库连接
// var db *gorm.DB

// func initDB() (err error) {
// 	username := "root"   //账号
// 	password := "123456" //密码
// 	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
// 	port := 3306         //数据库端口
// 	Dbname := "go_test"  //数据库名
// 	timeout := "10s"     //连接超时，10秒
// 	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
// 	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
// 	db, err := gorm.Open("mysql", dsn)
// 	if err != nil {
// 		panic("连接数据库失败, error=" + err.Error())
// 	}
// 	//延时关闭数据库连接
// 	defer db.Close()
// }
var db *sqlx.DB

func initDB() (err error) {
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "go_test"  //数据库名
	timeout := "10s"     //连接超时，10秒
	// 拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

// 查数据
func queryAllBook() (bookList []*Book, err error) {

	sqlStr := "SELECT  id,title,price FROM book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询所有书籍信息失败！")
		return
	}
	return
}

//  查询单条数据
func queryBookByID(id int64) (book Book, err error) {

	// sqlStr := "SELECT  id,title,price FROM book WHERE id=?"
	sqlStr := "SELECT book.id, book.price, book.title, publisher.province, publisher.city FROM book JOIN publisher ON book.publisher_id = publisher.id  WHERE book.id = ?;"

	err = db.Get(&book, sqlStr, id)
	if err != nil {
		fmt.Println("查询书籍信息失败！")
		return
	}
	fmt.Printf("%#v\n", book)
	return
}

// func queryBookByID(id int64) (book *Book, err error) {
// 	book = &Book{}
// 	sqlStr := "SELECT  id,title,price FROM book WHERE id=?"
// 	err = db.Get(book, sqlStr, id)
// 	if err != nil {
// 		fmt.Println("查询书籍信息失败！")
// 		return
// 	}
// 	return
// }

//  插入数据
func insertBook(title string, price float64) (err error) {
	sqlStr := "INSERT INTO  book (title, price)  VALUES (?,?);"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入数据失败", err)
		return
	}
	return
}

//  删除数据
func deleteBook(id int64) (err error) {
	sqlStr := "DELETE  FROM  book WHERE id= ?;"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除 书籍失败", err)
		return
	}
	return
}

//  day 11
func editBook(title string, price float64, id int64) (err error) {
	sqlStr := "UPDATE book SET title=?, price=? WHERE id = ?;"
	_, err = db.Exec(sqlStr, title, price, id)
	if err != nil {
		fmt.Println("更新书籍失败", err)
		return
	}
	return
}
