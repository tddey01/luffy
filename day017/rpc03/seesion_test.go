package rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

//// 测试 文件
//func TestSession_ReadWrite(t *testing.T) {
//	// 定义监听ip地址和端口
//	addr := "127.0.0.1:8000"
//	// 定义传输数据
//	my_data := "hello"
//	// 等待组 同步问题
//	wg := sync.WaitGroup{}
//	// 携程1个读 1个写
//	wg.Add(2)
//	// 写书进入携程
//	go func() {
//		defer wg.Done()
//		//	 创建tcp连接
//		lis, err := net.Listen("tcp", addr)
//		if err != nil {
//			t.Fatal(err)
//		}
//		conn, _ := lis.Accept()
//		s := Session{conn: conn}
//		//写数据
//		err = s.Write([]byte(my_data))
//		if err != nil {
//			t.Fatal(err)
//		}
//	}()
//	//	 读数据携程
//	go func() {
//		defer wg.Done()
//		conn, err := net.Dial("tcp", addr)
//		if err != nil {
//			t.Fatal(err)
//		}
//		s := Session{conn: conn}
//		data, err := s.Read()
//		if err != nil {
//			t.Fatal(err)
//		}
//		if string(data) != my_data {
//			t.Fatal(err)
//		}
//		fmt.Println(string(data))
//	}()
//	wg.Wait()
//}

func TestSession_ReadWriter(t *testing.T) {
	// 定义地址
	addr := "127.0.0.1:8000"
	my_data := "hello"
	// 等待组定义
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 写数据的协程
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, _ := lis.Accept()
		s := Session{conn: conn}
		err = s.Write([]byte(my_data))
		if err != nil {
			t.Fatal(err)
		}
	}()
	//读数据的协程
	go func() {
		defer wg.Done()
		conn, err := net.DialTimeout("tcp", addr, time.Second*5)
		//conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s := Session{conn: conn}
		data, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		// 最后一层校验
		if string(data) != my_data {
			t.Fatal(err)
		}
		fmt.Println(string(data))
	}()
	wg.Wait()
}
