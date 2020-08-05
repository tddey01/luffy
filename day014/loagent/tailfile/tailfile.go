package tailfile

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"github.com/tddey01/luffy/day014/loagent/kafka"
)

// tail 相关
//var (
//	confChan chan []common.CollectEntry
//)

type TailTask struct {
	path   string
	topic  string
	tObj   *tail.Tail
	ctx    context.Context
	cancel context.CancelFunc
}

func newTailTask(path, topic string) *TailTask {
	ctx, cancel := context.WithCancel(context.Background())
	tt := &TailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
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
		select {
		case <-t.ctx.Done(): // 只要调用t.cancel() 就会收到信号
			logrus.Infof("path:%s is stopping...", t.path)
			return
			// 循环读取数据
		case line, ok := <-t.tObj.Lines: // chan tail.Line
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
}
