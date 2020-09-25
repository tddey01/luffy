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

```
资源均来源于网络，下载将默认同意免责协议：
http://vmv.re/Erwyq

Workstation 16 Pro for Windows :
http://vmv.re/h0MYu
Workstation 16 Pro for Linux : 
http://vmv.re/OjNpz
VMware Fusion Pro 12.0.0 for Mac :
http://vmv.re/F28N4
Se-rial K-ey: 
FC11K-00DE0-0800Z-04Z5E-MC8T6
ZF3R0-FHED2-M80TY-8QYGC-NPKYF
YF390-0HF8P-M81RQ-2DXQE-M2UT6
ZF71R-DMX85-08DQY-8YMNC-PPHV8

27 July 2020 更新
NSX-T 3.0.1.1
http://vmv.re/6BlB7 提取码：vras

更新日志：
http://vmv.re/vmrd

常用资源下载：
vSphere 7.0.0b ，包括ESXi vSphere vCenter 
http://vmv.re/h6d9I 提取码：vras
vSphere 7 注册机：
http://vmv.re/ndM7c 提取码：vras
NSX-T 3.0.1:
http://vmv.re/zYAeK  提取码：vras
Vmware资源大全,所有资料：
百度云群组（进去之后保存资料然后退出百度云群组，总共就200人的大小）：

http://vmv.re/vmresource
http://vmv.re/vmresource1
http://vmv.re/vmresource2

 VMware Cloud Foundation 4.0.0   
链接：http://share.weiyun.com/YTCGS9Gk   密码：3dwn8x

备份站点：
Storage重新上线，暂时只支持大陆以外IP的请求。
vmware资源大全：
http://vmv.re/backvmrd
vSphere 7 ：
http://vmv.re/backv7d
vSphere 7 注册机：
http://vmv.re/backv7key
统一密码：new.llycloud.com
```
