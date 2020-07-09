package main

import (
	"fmt"
)

func defer_call() {
	defer func() { fmt.Println("print before recovery") }()
	defer func() { recover() }() // 正确
	defer func() { fmt.Println("print after recovery") }()
	panic("panic info")
}

func main() {
	defer_call()
}
