package main

import (
	"fmt"
)

// 打印印九九乘法表
func main() {
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d * %d = %d  ", j, i, j*i)
	// 	}
	// 	fmt.Println()
	// }
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, j*i)
		}
		fmt.Println()
	}
}
