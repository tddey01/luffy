package main

import (
	"fmt"
	"time"
)

// 内置的time包

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
	// timeObj := time.Unix(timestamp, 360*1000000) // 将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     // 年
	month := timeObj.Month()   // 月
	day := timeObj.Day()       // 日
	Hour := timeObj.Hour()     // 小时
	minute := timeObj.Minute() // 分钟
	second := timeObj.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n", year, month, day, Hour, minute, second)
}

func tickDemo() {
	ticker := time.Tick(time.Second * 2) // 定义一个1秒时间间隔的定时器

	for i := range ticker {
		fmt.Println(i) // 每秒都会执行的任务
	}
}

// 时间格式化
func formtDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分
	fmt.Println(now.Format("2006-01-02 15:04.999"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("2006年01月02日"))
	fmt.Println(now.Format("2006年01月02日 15点04分"))
}
func main() {
	// time.Time struct
	now := time.Now()
	fmt.Printf("%#v\n", now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	timestampDemo2(1592375005)
	//  定时器 执行
	// tickDemo()

	// 时间格式化
	// 日期格式化 时间格式化的数据 ==>> 字符串格式化
	formtDemo()
}
