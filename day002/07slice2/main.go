package main

import "fmt"

func main() {
	var a = []int{}
	fmt.Printf("a %v len:%d cap:%d ptr:%p\n",a, len(a), cap(a),a)
	a = append(a, 1)
	fmt.Printf("a %v len:%d cap:%d ptr:%p\n",a, len(a), cap(a),a)
	a = append(a, 2)
	fmt.Printf("a %v len:%d cap:%d ptr:%p\n",a, len(a), cap(a),a)
	a = append(a, 3)
	fmt.Printf("a %v len:%d cap:%d ptr:%p\n",a, len(a), cap(a),a)
}
