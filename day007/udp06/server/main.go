package main

import (
	"fmt"
	"net"
)

func process(listener *net.UDPConn) {
	defer listener.Close()
	// 循环接收数据
	for {
		var buf [1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接收消息失败 err:", err)
			return
		}
		fmt.Printf("接收到消息来自%v的消息，%v\n", addr, string(buf[:n]))
		// 回复消息
		n, err = listener.WriteToUDP([]byte("滚"), addr)
		if err != nil {
			fmt.Println("回复消息失败了", err)
			return
		}
	}
}

//  UDP server端
func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("listen :8080, 监听失败", err)
		return
	}
	process(listener)
}
