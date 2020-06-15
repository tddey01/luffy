package main

import "fmt"

//  函数是谁都可以调用
// 方法 就是某个具体类型才能调用的方法
type people struct {
	name   string
	gender string
}

// 函数指定接收者之后的方法
//  go 语言中约点成俗不用this也不用self， 而是使用后面类型首字母的小写
func (p *people) dream() {
	p.gender = "男"
	fmt.Printf("%s 我的梦想不用上班也有钱拿！ %s\n", p.name, p.gender)
}

func main() {
	var haojie = people{
		name:   "豪杰",
		gender: "爷们",
	}
	haojie.dream()
	(&haojie).dream()
	fmt.Println(haojie.gender)
}
