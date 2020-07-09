package main

import (
	"fmt"
	"runtime"
	"strconv"
)

//代码是否出发异常
//  会
func main() {
	runtime.GOMAXPROCS(1)
	for {
		int_chan := make(chan int, 1)
		string_chan := make(chan string, 1)
		int_chan <- 1
		string_chan <- "hehe"
		select {
		case value := <-int_chan:
			// fmt.Println(value + "out put")
			fmt.Println(strconv.Itoa(value) + "out put") // 转换
		case value := <-string_chan:
			fmt.Println(value + "out put")
		}
	}
}
