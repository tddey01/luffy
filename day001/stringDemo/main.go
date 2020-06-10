package main

import (
	"fmt"
)

// 字符串反转操作
func main() {
	// 第一种方法
	s1 := "hello"
	// s2 := "elloh"
	byterArry := []byte(s1) // [h e l l o]
	s2 := ""
	for i := len(byterArry) - 1; i >= 0; i-- {
		// i 是4 3 2 1
		// byterArry[i] o l l e h [字符串]
		s2 = s2 + string(byterArry[i])
	}
	fmt.Println(s2)
	// 方法二
	length := len(byterArry)
	for i := 0; i < length/2; i++ {
		byterArry[i], byterArry[length-1-i] = byterArry[length-1-i], byterArry[i]
	}
	fmt.Println(string(byterArry))
}
