package main

import "fmt"

// defer
func f1() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println("北京")
}
func main() {
	f1()
}
