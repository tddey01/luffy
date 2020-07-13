package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB  当你不知道变量名首字母是不是应该大写还是小写， 那你就用小写
// 保证最小暴露原则
var db *sqlx.DB

//  初始化数据库
func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test??charset=utf8&parseTime=true&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	// 连接成功
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

// 创建用户数据函数
func createUser(username, password string) error {
	sqlStr := "insert into user(usernma,password) values(?,?)"
	_, err := db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Println("插入失败", err)
		return err
	}
	return nil
}
