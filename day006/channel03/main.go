package main

import "fmt"

// 通道 接收值时判断通道是否关闭

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 不像 打开的文件必须在代码中显示的关闭 channel可与被垃圾回收机制回收
}

func main() {
	var ch1 = make(chan int, 100)
	go send(ch1)
	// // 利用for循环通道ch1中接收值
	// for {
	// 	ret, ok := <-ch1 // 使用value ，ok:=<-ch1 取值方式，当通道关闭的时候， ok:=false
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(ret)
	// }

	for ret := range ch1 { // 利用for range 循环通道取值
		fmt.Println(ret)
	}
}
