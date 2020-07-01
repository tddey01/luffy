package main

import (
	"fmt"
	"net" // 跟网络相关的代码
)

// 1 监听端口
// 2 接收客户端请求建立连接
// 3 创建goroutine处理连接

//  单独处理连接的函数
func process(conn net.Conn) {
	//  从连接中接收数据
	// var buf [1024]byte
	buf := make([]byte, 1024)
	n, err := conn.Read(buf[:]) // n表示读了多少数据
	if err != nil {
		fmt.Println("接收客户端发来的消息失败了，err:", err)
		return
	}
	fmt.Println("接收客户端发来的消息：", string(buf[:n]))
}

func main() {
	// 1 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen :8080, 监听失败", err)
	}
	defer listener.Close() // 程序退出时释放20000
	// 2 接收客户端请求建立连接
	for {
		conn, err := listener.Accept() // 如果没有客户端连接就阻塞，一直在等待
		if err != nil {
			fmt.Println("建立连接失败", err)
			continue
		}
		// 3 创建goroutine处理连接
		go process(conn)
	}
}
