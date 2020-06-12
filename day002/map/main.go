package main

import (
	"fmt"
	"strings"
)

func main() {
	// // 光声明map类型  但是没有初始化 a就是一个初始值
	// var a map[string]int
	// fmt.Println(a == nil)
	// //  map的初始化
	// a = make(map[string]int, 8)
	// fmt.Println(a == nil)
	// //  map中添加键值对
	// a["沙河"] = 100
	// a["上海"] = 200
	// fmt.Println(a)
	// fmt.Printf("type:%T\n", a)

	// // 声明map的同时完成初始化
	// b := map[int]bool{
	// 	1: true,
	// 	2: false,
	// }
	// fmt.Println(b)
	// fmt.Printf("b:%#v\n", b)
	// fmt.Printf("type:%T\n", b)

	// var c map[int]int
	// c = make(map[int]int, 2)
	// c[100] = 200 // c这个map没有初始化不能直接操作
	// fmt.Println(c)

	// // 判断某个键存不存在
	// var scoreMap = make(map[string]int, 0)
	// scoreMap["沙河"] = 100
	// scoreMap["上海"] = 200

	// // 判断 张二狗子 在不在scoreMap中
	// // value, ok := scoreMap["张二狗子"]
	// value, ok := scoreMap["上海"]
	// fmt.Println(value, ok)

	// if ok {
	// 	fmt.Println("张二狗在scoreMap", value)
	// } else {
	// 	fmt.Println("查无此人")
	// }

	// //  遍历
	// var scoreMap = make(map[string]int, 8)
	// scoreMap["张二狗子"] = 100
	// scoreMap["上海"] = 200
	// // for range
	// //  map是无序，键值对混合添加的顺序无关
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }
	// //  只遍历map中的key
	// for k := range scoreMap {
	// 	fmt.Println(k)
	// }

	// //  只遍历map中的value
	// for _, v := range scoreMap {
	// 	fmt.Println(v)
	// }

	// //  删除 上海 这个键值对
	// var scoreMap = make(map[string]int, 8)
	// scoreMap["张二狗子"] = 100
	// scoreMap["上海"] = 200
	// delete(scoreMap, "上海")
	// fmt.Println(scoreMap)

	// //  按照某个固定顺序遍历map
	// var scoreMap = make(map[string]int, 100)
	// // 添加50个键值对
	// for i := 0; i < 50; i++ {
	// 	key := fmt.Sprintf("stu%02d", i)
	// 	value := rand.Intn(100) // 0到99的随机整数
	// 	scoreMap[key] = value
	// }

	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }
	// //  按照key从小大的顺序遍历scoreMap
	// // 1 先去所有的key 存放到切片中
	// keys := make([]string, 0, 100)
	// for k := range scoreMap {
	// 	keys = append(keys, k)
	// }
	// // 2.对key做排序
	// sort.Strings(keys) //keys目前是一个有序的切片
	// // 3.排序后的key对scoreMap排序
	// for _, key := range keys {
	// 	fmt.Println(key, scoreMap[key])
	// }

	// // 元素类型为map的切片
	// var mapSlice = make([]map[string]int, 8, 8) // 只是完成切片初始化
	// fmt.Println(mapSlice[0] == nil)
	// //  还需要完成内部map元素的初始化
	// mapSlice[0] = make(map[string]int, 8) // 完成map的初始化

	// mapSlice[0]["沙河"] = 100
	// mapSlice[0]["北京"] = 200
	// fmt.Println(mapSlice)

	// // 值为切片的map
	// var sliceMap = make(map[string][]int, 8) //值完成了map的初始化
	// v, ok := sliceMap["中国"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	// slieceMap中没有重做这个键值
	// 	sliceMap["中国"] = make([]int, 8) // 完成了对切片初始化
	// 	sliceMap["中国"][0] = 100
	// 	sliceMap["中国"][1] = 200
	// 	sliceMap["中国"][2] = 300
	// }
	// // 遍历sliceMap
	// for k, v1 := range sliceMap {
	// 	fmt.Println(k, v1)
	// }

	//  统计一个字符串中每个单词出现次数
	// "how do you do" 中每个单词出现次数
	// 0 定义一个map[string]int
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	// 1 字符串中都有哪些单词
	words := strings.Split(s, " ")
	// 2 遍历单词做统计
	for _, word := range words {
		v, ok := wordCount[word]
		if ok {
			// map中有这个单词的统计记录
			wordCount[word] = v + 1
		} else {
			// map中没有单词统计记录
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}

}
