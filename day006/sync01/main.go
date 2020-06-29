package main

import (
	"fmt"
	"sync"
)

//sync
var x int64 // 定义全局变量x
var wg sync.WaitGroup

//  定义一个互斥锁
var lock sync.Mutex

//  定义一个函数 对全局的变量x做累加的操作
func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 上锁
		//  1 从内存中找到x值 0  500
		//  2 执行一个+1操作
		//  3 把结果赋值x 写到内存中
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
