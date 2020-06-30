package split

import (
	"fmt"
	"strings"
)

// 定义一个切割字符串的包

// Split 用sep分数
//  a:b:c ： --> [a b c]
func Split(s, sep string) (result []string) {
	index := strings.Index(s, sep)
	for index >= 0 {
		result = append(result, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func add(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}
