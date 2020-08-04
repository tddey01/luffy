package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"172.16.0.6:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed err:%v", err)
		return
	}
	fmt.Println("connect to etcd successfull!")

	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	str := `[
    {
        "path":"/Users/access/Projects/go_code/src/github.com/tddey01/luffy/day014/loagent/logs/s4.log",
        "topic":"s4_log"
    },
    {
        "path":"/Users/access/Projects/go_code/src/github.com/tddey01/luffy/day014/loagent/logs/web.log",
        "topic":"web_log"
    }
  ]`
	_, err = cli.Put(ctx, "collect_log_conf", str)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed err:%v\n", err)
		return
	}

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "collect_log_conf")

	if err != nil {
		fmt.Printf("get from etcd failed , err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	cancel()
}

// switch
// func main() {
// 	cli, err := clientv3.New(clientv3.Config{
// 		Endpoints:   []string{"127.0.01:2379"},
// 		DialTimeout: time.Second * 5,
// 	})
// 	if err != nil {
// 		fmt.Printf("connect to etcd failed err:%v", err)
// 		return
// 	}
// 	defer cli.Close()

// 	// watch
// 	watchCh := cli.Watch(context.Background(), "s4")
// 	for wresp:= range watchCh {
// 		for _, evt := range wresp.Events {
// 			fmt.Printf("type :%s  key:%s value:vs\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
// 		}
// 	}
// }
