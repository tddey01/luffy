package main

import "fmt"

// 切片
func main() {
	var a = [3]int{1, 2, 3} // 数组
	// 声明前排的方式1： 直接声明
	var b = []int{1, 2, 3} // 切片
	fmt.Println(a, b)
	fmt.Printf("a:%T b:%T\n", a, b)

	fmt.Println(b[1])

	// // 声明前排的方式2： 从数组得到切片
	// var c []int
	// c = a[0:2] // c=a[:] // 从数组的开始到结束
	// fmt.Printf("c:%T\n", c)
	// fmt.Println(c)

	// //  冒号切 左包含 右不包含
	// d := a[:2] // 从开始切刀索引是2(不包含)
	// e := a[:1] // 从索引为1开始切到最后
	// fmt.Println(d)
	// fmt.Println(e)

	//  切片的大小(目前元素的数量)
	fmt.Println(len(b))
	// 容量（底层数组最大能多少元素)
	x := [...]string{"北京", "上海", "深圳", "广州", "成都", "杭州", "西安", "重庆"}
	y := x[1:4]
	fmt.Println(y)
}
