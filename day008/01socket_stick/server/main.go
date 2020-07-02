package main

import (
	"bufio"
	"fmt"
	"net"
)

// 粘包现象 服务端
// socket_stick/server/main.go

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	buf := make([]byte, 1024)
	n, err := reader.Read(buf[:])
	if err != nil {
		fmt.Println("接收客户端发来的消息失败了，err:", err)
		return
	}
	fmt.Println("接收客户端发来的消息：", string(buf[:n]))
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen :8080, 监听失败", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("建立连接失败", err)
			continue
		}

		go process(conn)
	}
}
