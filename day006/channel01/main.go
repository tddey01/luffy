package main

import "fmt"

// channel
func main() {
	// 定义一个ch1变量
	//  是一个channel类型，
	// 这个channel内部传递数据是init类型
	var ch1 chan int
	var ch2 chan string
	// channel 是引用类型
	fmt.Println("ch1:", ch1)
	fmt.Println("ch2:", ch2)
	//  make 函数初始化（分配内存) slice  map channel
	ch3 := make(chan int, 1)
	//  通道操作 发送 接收 关闭
	//  发送和接收都用一个符号  <-
	ch3 <- 10    // 把10发送到ch3中
	ret := <-ch3 // 从ch3中接收值， 保存到变量ret中
	fmt.Println(ret)
	ch3 <- 20
	ret = <-ch3
	fmt.Println(ret)

	// 关闭
	close(ch3)

	//  1 关闭的通道 在接收， 能取到对应的零值
	ret2 := <-ch3
	fmt.Println(ret2)
	//  2往一个关闭通道中发送值， 会有panic错误
	// ch3 <- 30
	// 关闭一个已经关闭的通道  回应发panic
	close(ch3)
}
