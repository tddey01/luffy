package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	pb "github.com/tddey01/luffy/day017/rpc05/proto"
)

// 声明结构体
type Hello struct{}

// 实现接口方法
func (c *Hello) Info(ctx context.Context, req *pb.InfoRequest, rep *pb.InfoResponse) error {
	rep.Msg = "你好" + req.Name
	return nil
}

func main() {
	// 得到一个微服务实例
	server := micro.NewService(
		// 设置微服务名字。用来做访问用的
		micro.Name("hello"),
	)
	// 初始化
	server.Init()
	// 服务注册
	err := pb.RegisterHelloHandler(server.Server(), new(Hello))
	if err != nil {
		fmt.Println("服务注册失败", err)
	}
	// 启动微服务
	err = server.Run()
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
