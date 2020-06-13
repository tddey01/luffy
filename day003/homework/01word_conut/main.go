package main

import (
	"fmt"
	"strings"
)

//  写一个程序
// 统计一个字符串中每个单词出现的次数
// 比如： “how do you do ”how=1 do=2 you=1
func main() {
	s := "how do you do are you i am fine thank you"
	// 1 看下字符串中都有哪些单词(用空格分割字符串切片)
	wordSlice := strings.Split(s, " ")
	fmt.Println(wordSlice)
	// 2 挨个数一数单词
	wordMap := make(map[string]int, len(wordSlice))
	// 3 找个适合的数据类型把结果存起来
	// 3.1 找个适合的数据
	for _, value := range wordSlice {
		// fmt.Println(value)
		// 3.2.1 如果这个单词不在map中就添加兼职对，让次数为1
		v, ok := wordMap[value]
		if ok {
			wordMap[value] = v + 1
		} else {
			wordMap[value] = 1
		}

	}

	// 遍历map
	for k, v := range wordMap {
		fmt.Println(k, v)
	}
}
