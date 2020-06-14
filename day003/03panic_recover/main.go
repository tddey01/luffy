package main

import "fmt"

//painc 错误
func main() {
	defer func() {
		//ercover
		err := recover() // 尝试将函数从当前的异常状态恢复过来
		fmt.Println("recover 捕获到异常", err)
	}()
	var a []int
	a[0] = 100 // panic
	fmt.Println("这是main函数")
}
