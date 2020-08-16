package main

// log transfer
// 从kafka消费日志数据  写入es
import (
	"fmt"

	"github.com/tddey01/luffy/day016/log_transfer/es"

	"github.com/tddey01/luffy/day016/log_transfer/kafka"

	"github.com/go-ini/ini"
	"github.com/tddey01/luffy/day016/log_transfer/model"
)

func main() {
	var cfg = new(model.Config)
	err := ini.MapTo(cfg, "./conf/.editorconfig")
	if err != nil {
		fmt.Printf("load config  failed err:%v\n", err)
		panic(err)
	}
	fmt.Println("load config successfull....")
	// 连接kafka
	err = kafka.Init()
	if err != nil {
		fmt.Printf("connect to kafka failed err:%v\n", err)
		return
	}
	fmt.Println("connect to kafka successfull....")
	// 连接ES
	err = es.Init()
	if err != nil {
		fmt.Printf("connect to es failed err:%v\n", err)
		return
	}
	fmt.Println("connect to es successfull....")

}
