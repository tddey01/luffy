package main

import (
	"fmt"
)

// goto
func main() {
	flag := false
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				flag = true
				// 将i等于2 j等于2 跳出两层for循环
				break
			}
			fmt.Printf("%d--%d\n", i, j)
		}
		// 通过标志位来判断是否跳出外层for循环
		if flag {
			break
		}
	}
	fmt.Println("两层否循环结束")

	// for i := 0; i < 5; i++ {
	// 	if i == 2 {
	// 		// break
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	gotoDemo()
}

// goto
func gotoDemo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 退出设置标签
				goto breakTag
			}
			fmt.Printf("%d--%d\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}
