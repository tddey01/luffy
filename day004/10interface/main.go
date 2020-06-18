package main

import "fmt"

// 类型断言
// type nullInterface interface{}  // 太繁琐
type Cat struct{}

func showType(x interface{}) {
	//  因为我这个函数可以接受任何任意类型的变量
	// 类型断言
	v1, ok := x.(int)
	if !ok {
		// 说明猜错了
		fmt.Println("不是int", v1)
	} else {
		fmt.Println("x 就是一个int类型", v1)
	}
	v2, ok := x.(string)
	if !ok {
		// 说明猜错了
		fmt.Println("不是string")

	} else {
		fmt.Println("x 就是一个int类型", v2)
	}
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string value is %v\n", v)
	case int:
		fmt.Printf("x is a int value is %v\n", v)
	case bool:
		fmt.Printf("x is bool value is %v\n", v)
	case Cat:
		fmt.Printf("x is a Cat  struct value is %v\n", v)
	case *string:
		fmt.Printf("x is sa string poneinter value is %v\n", v)
	default:
		fmt.Println("unsupport type!")
	}

}

func main() {
	// var x interface{}
	// x = 100

	// showType(100)
	// showType("哈哈")

	//  switch 类型断言
	justifyType(100)

	justifyType(Cat{})
	justifyType("哈哈")
	s := "testhaha"
	justifyType(&s)

}
