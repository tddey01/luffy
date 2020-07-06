package main

import (
	// 只用到了他的这个包里面的init()
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// database/sql 连接MySQL示例代码
func main() {
	// dsn:"user:password@tcp(ip:port)/databasename"
	//   "dsn": "root:kAaSWE$%@tcp(192.168.192.214)/newdata1?charset=utf8&parseTime=true&loc=Local",
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8&parseTime=true&loc=Local"
	// 调用标准库中的方法
	// 前提是要注册对应数据库的驱动
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql failed, err:%v\n", err)
		return
	}
	// 尝试链接一下数据库，校验用户名密码是否正确...
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect MySQL failed, err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
}
