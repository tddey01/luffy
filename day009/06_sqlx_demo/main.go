package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB 全局数据库连接对象（内置连接池的）
var DB *sqlx.DB

// User user表对应的结构体
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	return nil
}

// 查询单条
func queryRowDemo() {
	sqlStr := "select id,name from users where id=?"
	var user User
	err := DB.Get(&user, sqlStr, 1) // get获取一条
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", user)
}

// 查询多行
func queryMultiDemo() {
	sqlStr := "select id,name,age from users where id >?"
	var users []User
	err := DB.Select(&users, sqlStr, 0) // select获取多条
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("user:%#v\n", user)
	}
}

//  事务
func transdemo() {
	tx, err := DB.Beginx() // 开始事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trnas failed, err:%v\n", err)
		return
	}
	sql1 := "update users set age=age+? where id=?"
	tx.MustExec(sql1, 2, 1) //名字带有Must的一般表示出错就panic
	tx.MustExec(sql1, 2, 2) //名字带有Must的一般表示出错就panic
	err = tx.Commit()
	if err != nil {
		fmt.Printf("commitfailederr:%v\n", err)
	}
	fmt.Println("两条数据更新成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}

	queryRowDemo()
	queryMultiDemo()
	transdemo()
}
