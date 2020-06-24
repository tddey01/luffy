package main

import (
	"fmt"
	"sync"
)

//  启动goroutine

var wg sync.WaitGroup

func hello() {
	fmt.Println("Hllo 北京")
	wg.Done()
}

func main() {
	defer fmt.Println("哈哈哈")
	wg.Add(10) // 计数器 10
	for i := 0; i < 10; i++ {
		go hello() // 创建一个goroutine  在新的goroutine中执行hello函数
	}

	fmt.Println("hello main func.")
	// time.Sleep(time.Second)
	//  等待hello执行完， 《执行hello函数的那个goroutine执行完》
	wg.Wait() // 阻塞， 一直等待所有的goroutine结束
}
