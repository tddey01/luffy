package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tddey01/luffy/day014/loagent/common"
	"github.com/tddey01/luffy/day014/loagent/tailfile"
	"go.etcd.io/etcd/clientv3"
)

// etc 相关操作

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
func GetConf(key string) (collectEntryList []common.CollectEntry, err error) {
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
	if err != nil {
		logrus.Errorf("json unmarshi failed , err:%v\n", err)
		return
	}
	return
}

// 监控etcd中配置变化收集项的函数
func WatchConf(key string) {
	for {
		watchCh := client.Watch(context.Background(), key)
		var newConf []common.CollectEntry
		for wresp := range watchCh {
			logrus.Info("get new conf from etcd!")
			for _, evt := range wresp.Events {
				//fmt.Printf("type :%s  key:%s  value%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					logrus.Errorf("WatchConf json unmarshal new conf failed err:%v\n", err)
					continue
				}
				//	 告诉 tailfile 这个模块 应该启用新的配置了
				tailfile.SendNewConf(newConf) // 没有接收 就是阻塞了 暂停状态
			}
		}
	}
}
