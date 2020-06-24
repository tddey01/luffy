package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()

	for i := 1; i < 10; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("8:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(8) // 设置我的go程序只是用一个逻辑核心 m:n 中设置n为1
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
	// time.Sleep(time.Second)
}
