package main

import (
	"fmt"
)

func main() {
	s := "hello中国"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // 默认安装ascii码去打印
		fmt.Printf("%c\n", s[i])
	}

	for k, v := range s {
		fmt.Printf("%d %c\n", k, v)
		fmt.Println(k, v)
	}
}
