package main

import "fmt"

func main() {
	// 定义一个数组
	a := [...]int{1, 3, 5, 7, 9, 11, 13}
	// 基于数组得到一个切片
	b := a[:]
	// 修改切片中的第一个元素为1000
	b[0] = 100
	//  打印数组中第一个元素的值
	fmt.Println(a[0])

	c := a[2:5]
	fmt.Println(c)      //[5 7 9]
	fmt.Println(len(c)) //3
	fmt.Println(cap(c)) //5
	fmt.Printf("c:%p\n", c)

	d := c[:5]
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	fmt.Printf("d:%p\n", d)

	e := d[2:]
	fmt.Println(e)      //[9 11 13]
	fmt.Println(len(e)) //3
	fmt.Println(cap(e)) //3
	fmt.Printf("d:%p\n", e)
	e = append(e, 100, 200, 300)
	fmt.Println(e)
	fmt.Println(len(e)) //3
	fmt.Println(cap(e)) //3
	fmt.Printf("d:%p\n", e)

	e[0] = 900
	fmt.Println(e)
	fmt.Println(len(e)) //6
	fmt.Println(cap(e)) //6

	// make 函数构造切片
	f := make([]int, 5, 10)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	// 通过len()函数获取切片长度
	fmt.Println(len(f))
	//  通过cap()函数获取切片的容量
	fmt.Println(cap(f))

	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}
