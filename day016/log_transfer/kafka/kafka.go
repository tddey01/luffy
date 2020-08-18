package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 初始化kafka
// 从kafka取出数据
v
func Init(address []string, topic string) (err error) {
	// 创建连接消费者
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	// 拿到消息指定topic 下面的所有的分区列表
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to list  partition, err:%v\n", err)
		return
	}

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的翻去的消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d , err :%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费者信息
		fmt.Println("start to consume...")
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
				logDataChan <- msg // 为了将同步流程异步化,所以将取出的日志数据先放到channel中
			}
		}(pc)
	}
	return
}
