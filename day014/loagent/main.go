package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"github.com/tddey01/luffy/day014/loagent/etcd"
	"github.com/tddey01/luffy/day014/loagent/kafka"
	"github.com/tddey01/luffy/day014/loagent/tailfile"
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

// 真正的业务逻辑
func run() (err error) {
	// TailObj --> log --> Client --> kafka
	for {
		line, ok := <-tailfile.TailObj.Lines // chan tail.Line
		if !ok {
			logrus.Warning("tail file close reopen, filename:%s\n", tailfile.TailObj.Filename)
			time.Sleep(time.Second) // 读取出错等一秒
			continue
		}
		//如果是空行 就跳过
		fmt.Printf("%#v\n", line.Text)
		if len(strings.Trim(line.Text, "\n\r")) == 0 {
			logrus.Info("空行")
			continue
		}
		// 利用通道将同步的代码改为异步的
		//fmt.Println("msg:", msg.Text)
		// 把读出来一行日志包装成kafka里面msg 类型，丢到通道中
		msg := &sarama.ProducerMessage{}

		msg.Value = sarama.StringEncoder(line.Text)

		//	 丢到管道中
		kafka.ToMsgChan(msg)
	}
}

func main() {
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
	err := ini.MapTo(configObj, "./conf/config.ini")
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
	logrus.Info("init  kafka succcess!")

	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		logrus.Errorf("init  etcd failed err:%v\n", err)
		return
	}
	logrus.Info("init  kafka succcess!")
	// 从etcd中拉去要收集日志的配置
	allConf, err := etcd.GetConf(configObj.EtcdConfig.CollectKey)
	if err != nil {
		logrus.Errorf("init  etcd failed allConf err:%v\n", err)
		return
	}
	fmt.Println(allConf)
	//	根据配置文件日志路径使用tail去收集日志
	err = tailfile.Init(configObj.CollectConfig.LogFilePath)
	if err != nil {
		logrus.Errorf("init  tailfile failed err:%v\n", err)
		return
	}
	logrus.Info("init  tailfile succcess!")

	//	把日志通过sarama发送kafka
	err = run()
	if err != nil {
		logrus.Errorf("run kafka  tailfile failed err:%v\n", err)
		return
	}
}
