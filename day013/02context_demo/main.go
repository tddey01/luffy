package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 初始化
func worker() {
	defer  wg.Done()
	//	 如何接收外部命令实现退出
	for {
		fmt.Println("worker")
		//time.Sleep(time.Second)
	}


}

func main(){
	wg.Add(1)
	go worker()
	// 如何优雅的实现结束 goroutine
	wg.Wait()
	fmt.Println("over")
}