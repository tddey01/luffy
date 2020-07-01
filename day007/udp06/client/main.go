package main

import (
	"fmt"
	"net"
)

//  UDP client
func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接失败server err", err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("约吗?"))
	if err != nil {
		fmt.Println("发送消息失败", err)
		return
	}
	// 接收消息
	buf := make([]byte, 1024)
	n, err = conn.Read(buf[:])
	if err != nil {
		fmt.Println("读取消息失败err", err)
		return
	}
	fmt.Println("收到回复", string(buf[:n]))
}
