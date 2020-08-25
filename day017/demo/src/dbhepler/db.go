package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tddey01/luffy/day017/demo/src/share/config"
)

//建表语句
var schema = `
CREATE TABLE  IF NOT EXISTS  user(
    id INT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(20) ,
    address VARCHAR(20),
    phone VARCHAR(15),
    PRIMARY KEY (id)
)
`

// 对应表的结构体
type User struct {
	Id      int32  `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
}

func main() {
	// 打开并连接数据库
	sqlx.Connect("mysql", config.MysqlDNS )
}
