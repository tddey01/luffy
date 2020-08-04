package tailfile

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"github.com/tddey01/luffy/day014/loagent/common"
	"github.com/tddey01/luffy/day014/loagent/kafka"
	"strings"
	"time"
)

// tail 相关
type TailTask struct {
	path  string
	topic string
	tObj  *tail.Tail
}

func newTailTask(path, topic string) *TailTask {
	tt := &TailTask{
		path:  path,
		topic: topic,
	}
	return tt
}

func (t *TailTask) Init() (err error) {
	cfg := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t.tObj, err = tail.TailFile(t.path, cfg)
	return
}

func (t *TailTask) run() {
	logrus.Infof("collect for path :%s is running...", t.path)
	//	 读取日志 发往kafka
	for {
		line, ok := <-t.tObj.Lines // chan tail.Line
		if !ok {
			logrus.Warning("tail file close reopen, path:%s\n", t.path)
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
		msg.Topic = t.topic // 每一个tailObj自己topic
		msg.Value = sarama.StringEncoder(line.Text)
		//	 丢到管道中
		kafka.ToMsgChan(msg)
	}
}

func Init(allConf []common.CollectEntry) (err error) {
	// allConf 里面存放若干个收集项
	// 针对每一个日志收集项创建一个对应的tailObj
	for _, conf := range allConf {
		tt := newTailTask(conf.Path, conf.Topic) // 创建了一个日志收集任务
		err = tt.Init() // 去打开日志文件准备读取
		if err != nil {
			logrus.Errorf("create tailObj  for path:%s failed err:%v", conf.Path, err)
			continue
		}
		logrus.Infof("create a tail task for path:%s successfull! ", conf.Path)
		//	起一个后台的goroutine 去收集日志
		go tt.run()
	}
	return
}
