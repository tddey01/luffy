package main

import (
	"fmt"
)

// 答应200到1000之间的的质数（素数 ）

func main() {
	for i := 200; i < 1000; i++ {
		// fmt.Println(i)
		flag := true
		// 判断i是否为质数 如果是就打印，如果不是不打印输出
		for j := 2; j < i; j++ {
			if i%j == 0 {
				// 不是值数
				flag = false
				break
			}
		}
		// 整个第二层的for循环没有被break说明是质数
		if flag {
			fmt.Printf("%d 是质数\n", i)
		}
	}
}
