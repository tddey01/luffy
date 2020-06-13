package main

import "fmt"

//  map
func main() {
	var n1 map[string]int
	n1 = make(map[string]int, 3) // 初花map
	n1["沙河"] = 1
	fmt.Println(n1)
}
