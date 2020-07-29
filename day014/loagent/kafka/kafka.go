package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// kafka 相关操作
var (
	Client sarama.SyncProducer
	MsgChan chan *sarama.ProducerMessage
)

func Init(address []string) (err error) {
	// 1. 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分区
	config.Producer.Return.Successes = true                   // 确认

	// 2. 连接kafka
	Client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		logrus.Error("kafka producer closed, err:", err)
		return
	}
//	初始化MsgChan
MsgChan = make(chan  *sarama.ProducerMessage, )
	return
}

