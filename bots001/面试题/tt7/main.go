package main

import (
	"fmt"
	"sync"
)

//  下面代码输出什么
func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("j", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
