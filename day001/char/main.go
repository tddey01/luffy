package main

import (
	"fmt"
)

// 字符串
func main() {
	s1 := "Golang"
	c1 := 'G' // ASCII码占用一个字节(8位 8bit)
	fmt.Println(s1, c1)
	s2 := "中国"
	c2 := '中' // UTF-8编码下一个中文占用3个字节
	fmt.Println(s2, c2)

	s3 := "hello沙河"
	fmt.Println(len(s3))

	// 遍历字符串
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%c\n", s3[i])
	}
	fmt.Println()
	// for range 循环 是按照rune类型去遍历
	for k, v := range s3 {
		fmt.Printf("%d %c\n", k, v)
	}

	// 修改字符串
	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	s5 := "big"
	// 将字符串强制转换成字节数组类型
	byterArry := []byte(s5)
	fmt.Println(byterArry)
	byterArry[0] = 'p'
	fmt.Println(byterArry)
	// 将字节数组强制转换成字符串
	s5 = string(byterArry)
	fmt.Println(s5)
}
