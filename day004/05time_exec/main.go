package main

import (
	"fmt"
	"time"
)

//  time 包练习
// 编写一个函数， 接收time类型的参数，函数内部将传进来的时间格式化输出2017/06/19 20:30:05 格式
//  编写程序统计一段代码的执行耗时时间 ， 单位精确到微妙
func printTime(t time.Time) {
	// "2017/06/19 20:30:05"
	//  yyy/m/d H:i:s
	timeStr := t.Format("2006/01/02 15:04:05")
	fmt.Println(timeStr)
}

func calcTime() {
	start := time.Now()
	startTimestamp := start.UnixNano() / 1000
	fmt.Println("钗头凤 红酥手 黄藤酒 满园春色宫墙柳")
	time.Sleep(time.Microsecond * 30)
	end := time.Now()
	endTimestamp := end.UnixNano() / 1000

	fmt.Println("耗费了时间", time.Since(start))
	fmt.Printf("耗费了%d微妙\n", endTimestamp-startTimestamp)
}
func main() {
	now := time.Now()
	printTime(now)
	calcTime()
}
