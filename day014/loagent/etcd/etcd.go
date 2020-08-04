package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// etc 相关操作
type collectEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

var (
	client *clientv3.Client
)

func Init(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed , err:%v", err)
		return
	}
	return
}

// 拉去日志收集的配置项目函数
func GetConf(key string) (collectEntryList []collectEntry,err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		logrus.Errorf("get conf from etcd by key:%v failed err:%v\n", key, err)
		return
	}
	if len(resp.Kvs) == 0 {
		logrus.Warning("get len:0 conf from etcd by key%s\n", key)
		return
	}
	ret := resp.Kvs[0]
	//ret.Value // json格式字符串
	fmt.Println(ret.Value)
	err = json.Unmarshal(ret.Value, &collectEntryList)
	if err !=nil{
		logrus.Errorf("json unmarshi failed , err:%v\n",err)
		return
	}
	return
}
