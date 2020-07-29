package tailfile

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

// tail 相关

var (
	TailObj *tail.Tail
)

func Init(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	// 打开文件开始读取数据
	TailObj, err = tail.TailFile(filename, config)
	if err != nil {
		logrus.Error("tailfile: create tailfObj for path:%s failed, err:%v\n", filename, err)
		return
	}
	return
}
