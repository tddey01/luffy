package main

import (
	"fmt"
)

//  下面代码输出啥
// methodA from parent
// MethodB from parent

type parent struct{}

func (p *parent) MethodB() {
	fmt.Println("MethodB from parent")
}

func (p *parent) MethodA() {
	fmt.Println("methodA from parent")
	p.MethodB()
}

type Child struct {
	parent
}

func (b *Child) MethodB() {
	fmt.Println("methodB crom child")
}
func main() {
	child := Child{}
	child.MethodA()
}
