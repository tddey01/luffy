package db

// 操作db
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 声明数据实例
var (
	db *sqlx.DB
)

func Init(mysqlDNS string) {
	// 获取连接
	db = sqlx.MustConnect("mysql", mysqlDNS)
	// 设置限制连接数
	db.SetMaxIdleConns(1)
	//最大打开的连接数
	db.SetMaxOpenConns(3)

}
