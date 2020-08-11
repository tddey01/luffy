package main

import (
	"fmt"

	"github.com/tddey01/luffy/day015/loagent/common"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"github.com/tddey01/luffy/day015/loagent/etcd"
	"github.com/tddey01/luffy/day015/loagent/kafka"
	"github.com/tddey01/luffy/day015/loagent/tailfile"
)

//  日志手机客户端
// 类似的开源项目 filebeat
// 收集指定目录下的日志文件， 发送到kafaka中

// 现在的进包
// 往kafka发送数据
// 使用tail读取日志文件

// Config 整个logent的配置结构体
type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
	EtcdConfig    `ini:"etcd"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	ChanSize int64  `ini:"chan_size"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

func run() {
	select {}
}

func main() {
	// -1 获取本地IP地址， 为后续去etcd获取配置文件打下基础
	ip, err := common.GetOuboundIP()
	if err != nil {
		logrus.Errorf("get ip failed ,err:%v", err)
		return
	}
	//	 初始化
	var configObj = new(Config)

	//	 读配置文件
	//cfg, err := ini.Load("./conf/config.ini")
	//if err != nil {
	//	logrus.Error("loal confnig failed, err:%v", err)
	//	return
	//}
	//kafkaAddr := cfg.Section("kafka").Key("address").String()
	//fmt.Println(kafkaAddr)
	err = ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		logrus.Errorf("loal confnig failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", configObj)

	//  初始化链接 kafka
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Errorf("init  kafka failed err:%v\n", err)
		return
	}
	logrus.Info("init  kafka succcessfull!")

	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		logrus.Errorf("init  etcd failed err:%v\n", err)
		return
	}
	logrus.Info("init  etcd succcessfull!")
	// 从etcd中拉去要收集日志的配置
	collcetKey := fmt.Sprintf(configObj.EtcdConfig.CollectKey, ip)
	allConf, err := etcd.GetConf(collcetKey)
	if err != nil {
		logrus.Errorf("init  etcd failed allConf err:%v\n", err)
		return
	}
	fmt.Println(allConf)
	// 派一个小弟  去监控etcd中， configObj.EtcdConfig.CollectKey 对应变化
	go etcd.WatchConf(collcetKey)
	//	根据配置文件日志路径使用tail去收集日志
	err = tailfile.Init(allConf) // 把从etcd中加载获取的配置项传到etcd中
	if err != nil {
		logrus.Errorf("init  tailfile failed err:%v\n", err)
		return
	}
	logrus.Info("init  tailfile succcessfull!")
	run()
}
