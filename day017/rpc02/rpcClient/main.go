package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 声明接收的参数结构体
type ArithRequest struct {
	A, B int
}

// 声明返回客户端参数结构体
type ArithResponse struct {
	// 成绩
	Pro int
	//	商
	Quo int
	//	余数
	Rem int
}

// 调用服务
func main() {
	//连接远程RPC
	//conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	req := ArithRequest{9, 10}
	var res ArithResponse
	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d*%d = %d\n", req.A, req.B, res.Pro)

	//	调用商
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d / %d 商= %d， 余数=%d\n", req.A, req.B, res.Quo, res.Rem)
}
