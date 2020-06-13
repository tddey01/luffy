package main

import "fmt"

func main() {
	// var a = [3][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// }
	// fmt.Println(a)
	b := [...]string{"北京", "上海", "深圳"}
	// 第一种
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
	fmt.Println()
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	// 第二种方式
	for index, value := range b {
		fmt.Println(index, value)
	}

	for index, _ := range b {
		fmt.Println(index)
	}

	for _, value := range b {
		fmt.Println(value)
	}
}
