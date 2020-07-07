package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// MYSQL 预处理
// DB 数据库连接句柄 全局变量
var Db *sql.DB

// User 用户结构体
type User struct {
	id   int
	name string
	age  string
}

func initDB(dsn string) (err error) {
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	Db.SetMaxIdleConns(50)
	return nil
}

func prepareDemo() {
	sqlStr := "insert into users (name,age) values (?,?)"
	stmt, err := Db.Prepare(sqlStr) // 把要执行的命令发给MYSQL服务器端 预处理
	if err != nil {
		fmt.Printf("prepare failed , err:%v\n", err)
		return
	}
	defer stmt.Close()
	//  之心循环的插入命令
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("stu%02d", i)
		stmt.Exec(name, i)
	}
}

// 预处理查询
func prepareQueryDemo() {
	sqlStr := "select id,name,age from users where id=?"
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepareSquery failed , err %v\n", err)
		return
	}
	defer stmt.Close()
	for i := 0; i < 10; i++ {
		rows, err := stmt.Query(i)
		if err != nil {
			fmt.Printf("query failed err%v\n", err)
			continue
		}
		defer rows.Close()
		var user User
		for rows.Next() {
			err := rows.Scan(&user.id, &user.name, &user.age)
			if err != nil {
				fmt.Printf("query failed err%v\n", err)
				return
			}
			fmt.Printf("user:%#v\n", user)
		}
	}
}
func main() {
	dsn := "root:123456@tcp(127.0.0.1)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Printf("init DB failed , err:%v", err)
		return
	}
	//  插入10条数据
	// prepareDemo()
	//  查询数据
	prepareQueryDemo()

}
