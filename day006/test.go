// package main

// import (
// 	"fmt"
// 	"os"
// )

// func getInput() {
// 	tmp := [1]byte{}
// 	os.Stdin.Read(tmp[:]) //标准输入获取值
// 	fmt.Println("123")

// }

// func main() {
// 	getInput()
// }

package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var m = make(map[string]int)
// 	fmt.Println(m["gdm"])
// 	m["gdm"] = 100
// 	fmt.Println(m["gdm"])
// }

func f1(a int) {
	fmt.Println(a)
}

func closer(x int) func() {
	return func() {
		f1(x)
	}
}

var onlyOne sync.Once

func main() {
	f := closer(10) // 利用闭包实现
	onlyOne.Do(f)
}
