package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 声明算数结构体
type Arith struct{}

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

// 乘法运算
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) (err error) {
	res.Pro = req.A * req.B
	return nil
}

// 商和余数
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) (err error) {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	// 商
	res.Quo = req.A / req.B
	// 余数
	res.Rem = req.A % req.B
	return nil
}

//func main() {
//	// 注册服务
//	rpc.Register(new(Arith))
//	// 采用http 作为rpc载体
//	rpc.HandleHTTP()
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		fmt.Println("启动服务失败")
//	}
//}
func main() {
	// 注册服务
	rpc.Register(new(Arith))
	// 采用http 作为rpc载体
	rpc.HandleHTTP()
	list, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("启动服务失败")
	}
	//	 循环监听服务
	for {
		conn, err := list.Accept()
		if err != nil {
			continue
		}
		// 携程
		go func(conn net.Conn) {
			fmt.Println("net client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
