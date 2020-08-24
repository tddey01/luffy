package main

import (
	"context"
	"fmt"

	pb "github.com/tddey01/luffy/day017/rpc04/gRPC"
	"google.golang.org/grpc"
)

func main() {
	// 创建与服务端连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常err:%v\n", err)
	}
	defer conn.Close()
	//	实例化RPC客户端
	client := pb.NewUserInfoServiceClient(conn)
	//	组装参数
	req := new(pb.UserRequest)
	req.Name = "zz"
	//	调用接口
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应结果异常 %v\n", err)
	}
	fmt.Println(resp)
}
