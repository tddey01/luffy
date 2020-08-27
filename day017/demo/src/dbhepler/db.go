package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/util/log"
	"github.com/tddey01/luffy/day017/demo/src/share/config"
	"math/rand"
	"time"
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
	sqlx.Connect("mysql", config.MysqlDNS)
}

// 按指定个数生成随机数
func GetRandomString(leng int) string {
	str := "0123456789abcdefghigkmnzxopqrstuvwyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
