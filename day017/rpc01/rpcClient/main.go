package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

// 调用服务
func main() {
	// 连接远程RPC服务
	rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	// 调用远程方法
	//	定义服务端传回来的计算结果的变量
	ret := 0

	// 求面积
	err = rp.Call("Rect.Area", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("面积", ret)
	// 2 周长
	err = rp.Call("Rect.Perimter", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长", ret)
}
