package main

import (
	"fmt"
	"math"
)

// 下面的max函数有什么问题
func max(a, b int64) int64 {
	return int64(math.Max(float64(a), float64(b)))
}

func main() {
	fmt.Println(max(1, 2))
	fmt.Println(max(math.MaxInt64-2, math.MaxInt64-1), math.MaxInt64-1)
}
