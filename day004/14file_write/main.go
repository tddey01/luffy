package main

import (
	"fmt"
	"os"
)

// 打开文件支持文件写入

func main() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("打开文件失败了", err)
		return
	}
	defer file.Close()
	str := "hello 沙河\n"
	file.Write([]byte("哈哈哈\n"))
	file.WriteString(str)

}
