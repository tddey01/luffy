package main

import "fmt"

//  连接操作
type student struct {
	name string
}

func (s *student) learn() *student {
	fmt.Printf("%s 热爱学习\n", s.name)
	return s
}

func (s *student) doHomework() {
	fmt.Printf("%s 热爱学习\n", s.name)
}

func main() {
	haojie := &student{"豪杰"}
	haojie.learn().doHomework()

	// ret := haojie.learn()
	// ret.doHomework()

}
