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
	// w代表大小写字母、数字、下划线
	reEmail = `\w+@w+\.\w+?`

	// 超链接正则
	// ？代表有没有s
	//  + 代表出1次或者多次
	// \s\S\ 各种字符
	//  +? 代表贪婪模式
	reLinke = `href="https?://[\s\S]+?"`

	//  手机号正则
	rePhone = `1[23456789]\d\s?\d{4}\s?\d{4}`

	// 身份证号
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))(0(1-9)|(1[012]))((0[1-9])|([12]\d)(3[01]))\d{3}[\dXx]`

	//  图片
	reImg =`https?://[^"]+?((jpg)|(png)|(jpeg)|(gif)|(bmp))`
)

// // GetEmail 获取邮箱
// func GetEmail() {
// 	resp, err := http.Get("https://tieba.baidu.com/p/601076813?red_tag=1573533731")
// 	HandleError(err, "http.Get url")
// 	defer resp.Body.Close()
// 	// 2 读取页面内容
// 	pageBytes, err := ioutil.ReadAll(resp.Body)
// 	HandleError(err, "ioutil.ReadAl")
// 	//字节转字符串
// 	pageStr := string(pageBytes)
// 	// fmt.Println(pageStr)

// 	//  3  过滤数据
// 	re := regexp.MustCompile(reQQEmail)
// 	//  -1 代表全部
// 	results := re.FindAllStringSubmatch(pageStr, -1)
// 	// fmt.Println(results)
// 	// 遍历
// 	for _, result := range results {
// 		fmt.Println("email", result[0])
// 		fmt.Println("qq", result[1])
// 	}
// }

// 对数组切片进行去重
func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// GetEmail 获取邮箱
func GetEmail(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println("email", result[0])
		fmt.Println("qq", result[1])
	}
}

// GetPageStr 根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAl")
	//字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

// HandleError 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func main() {
	//  爬取email
	// GetEmail()
	url := "https://tieba.baidu.com/p/601076813?red_tag=1573533731"
	// GetEmail(url)

	//  爬连接
	// url = "https://news.baidu.com"
	// GetLink(url)

	//  爬手机号
	// url = "https://www.zhaohaowang.com/"
	// GetPhone(url)

	//  身份证号
	// url = "https://henan.qq.com/a/20171107/069413.htm"
	// GetIdCard(url)

	//  爬图片
	url = "http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3"
	GetImg(url)
}

// GetLink 爬连接
func GetLink(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// GetPhone 爬取手机号
func GetPhone(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// GetIdCard 身份证号
func GetIdCard(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func GetImg(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}
