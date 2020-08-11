etcd clientv3 3.22 

```bash
# github.com/coreos/etcd/clientv3/balancer/resolver/endpoint
../../pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:114:78: undefined: resolver.BuildOption
../../pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:182:31: undefined: resolver.ResolveNowOption
# github.com/coreos/etcd/clientv3/balancer/picker
../../pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/picker/err.go:37:44: undefined: balancer.PickOptions
../../pkg/mod/github.com/coreos/etcd@v3.3.18+incompatible/clientv3/balancer/picker/roundrobin_balanced.go:55:54: undefined: balancer.PickOptions
```

## 解决方法

将grpc版本替换成v1.26.0版本
 1. 修改依赖为v1.26.0
```bash
go mod edit -require=google.golang.org/grpc@v1.26.0
```
下载v1.26.0版本的grpc
```bash
go get -u -x google.golang.org/grpc@v1.26.0
```
结果
```bash
connect to etcd successfull!
q1mx:dsb
```
## 案例代码
```bash
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
	_, err = cli.Put(ctx, "q1mx", "dsb")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed err:%v\n", err)
		return
	}
	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "q1mx")
	if err != nil {
		fmt.Printf("get from etcd failed , err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	cancel()
```