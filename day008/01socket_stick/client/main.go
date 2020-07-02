package main

import (
	"fmt"
	"net"
)

//  粘包现象 客户端
//  socket_stick/client/mian.go
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "hello,hello,how are you?"
		conn.Write([]byte(msg))
	}
}
