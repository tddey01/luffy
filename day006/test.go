package main

import (
	"fmt"
	"os"
)

func getInput() {
	tmp := [1]byte{}
	os.Stdin.Read(tmp[:]) //标准输入获取值
	fmt.Println("123")

}

func main() {
	getInput()
}
