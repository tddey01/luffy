package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// 爬邮箱

//  正则
var (
	reQQEmail = `(\d+)@qq.com`
)

// GetEmail 获取邮箱
func GetEmail() {
	resp, err := http.Get("https://tieba.baidu.com/p/601076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAl")
	//字节转字符串
	pageStr := string(pageBytes)
	// fmt.Println(pageStr)

	//  3  过滤数据
	re := regexp.MustCompile(reQQEmail)
	//  -1 代表全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	// fmt.Println(results)
	// 遍历
	for _, result := range results {
		fmt.Println("email", result[0])
	}
}

// HandleError 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	GetEmail()
}
