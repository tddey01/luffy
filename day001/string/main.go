package main

import (
	"fmt"
	"strings"
)

// 字符串操作
func main() {
	s1 := "alexdsb"
	fmt.Println(len(s1))

	// 字符串拼接
	s2 := "python"
	fmt.Println(s1 + s2)
	s3 := fmt.Sprintf("%s---%s", s1, s2)
	fmt.Println(s3)

	// 分割
	ret := strings.Split(s1, "x")
	fmt.Println(ret)

	// 判断是否包含
	ret1 := strings.Contains(s1, "sb")
	fmt.Println(ret1)

	// 判断前缀和后缀
	ret3 := strings.HasPrefix(s1, "alex")
	ret4 := strings.HasSuffix(s1, "sb")
	fmt.Println(ret3, ret4)

	// 求字符串的位置
	s4 := "applepen"
	ret5 := strings.Index(s4, "p")
	ret6 := strings.LastIndex(s4, "p")
	fmt.Println(ret5, ret6)

	//join
	a1 := []string{"Python", "PHP", "Javascript", "Ruby", "Golang"}
	fmt.Println(strings.Join(a1, "-"))
}
