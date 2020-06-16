package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

func main() {
	var stu1 = Student{
		ID:     1,
		Gender: "男",
		Name:   "李三",
	}
	//  序列化：把编程语言里面的数据转换成 JSON 格式化字符串
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("JOSN格式化错了", err)
	}
	fmt.Println(v)         // []byte
	fmt.Println(string(v)) // 把[]byte转成string
	fmt.Printf("%#v\n", string(v))

	// str := "{\"Id\":1,\"Gender\":\"男\",\"Name\":\"李三\"}"
	// // 反序列化：把满足JSON格式化字符串转成当前编程语言里面对象
	// var stu2 = &Student{}
	// json.Unmarshal([]byte(str), stu2)
	// fmt.Println(stu2)
	// fmt.Printf("%T\n", stu2)
}
