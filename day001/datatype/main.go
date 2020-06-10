package main

import (
	"fmt"
	"math"
)

func main() {
	// 十进制
	var a int = 10
	// 八进制  以0开头
	var b int = 077
	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Println(a, b, c)
	fmt.Printf("%b\n", a) // 1010  占位符%b表示二进制
	fmt.Printf("%o\n", b) // 77
	fmt.Printf("%x\n", c) // ff
	fmt.Printf("%X\n", c) // FF

	// 求c的变量内存地址
	fmt.Printf("%p\n", &c) //0xc0000b4008

	// 浮点数常量
	fmt.Println(math.MaxFloat64)  //1.7976931348623157e+308
	fmt.Println(math.MaxFloat32)  //3.4028234663852886e+38
	fmt.Printf("%f\n", math.Pi)   //3.141593
	fmt.Printf("%.2f\n", math.Pi) //3.14

	// 复数

	// 布尔值
	// 	Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。
	// 注意：
	// 布尔类型变量的默认值为false。
	// Go 语言中不允许将整型强制转换为布尔型.
	// 布尔型无法参与数值运算，也无法与其他类型进行转换。

	// 字符串
	// Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。 Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：

	// 字符串转义符
	fmt.Println("/usr/local/go/bin")
	var s1 = "单行文本"
	var s2 = `这个
	是 "这里不需要转义"
	多
	行
	文
	本`
	fmt.Println(s1, s2)
	// 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。

	
}
