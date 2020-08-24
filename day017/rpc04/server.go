package main

import (
	context "context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/tddey01/luffy/day017/rpc04/gRPC"
)

// 定义服务端实现约定的接口
type UserInfoService struct {
}

var u = UserInfoService{}

// 实现服务端实现的需要的接口
func (s *UserInfoService) GetUserInfo(cxt context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	// 在数据库查用户信息
	if name == "zz" {
		resp = &pb.UserResponse{
			Name: name,
			Age:  22,
			//切片字段
			Hobby: []string{"sing", "Run"},
		}
	}
	err = nil
	return
}

func main() {
	//   监听
	addr := "127.0.0.1:8080"
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("监听异常:", err)
	}

	fmt.Println("开始监听服务", addr)
	//	 实例化RPC
	s := grpc.NewServer()
	//	 在RPC上注册微服务
	// 要求接口累心变量
	pb.RegisterUserInfoServiceServer(s, &u)
	//	 启动RPC服务端
	s.Serve(conn)

}
