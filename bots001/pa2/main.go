package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	reImg = `https?://[^"]+?((jpg)|(png)|(jpeg)|(gif)|(bmp))`
)

//  获取页面数据
func myTest() {
	pageStr := GetPageStr("https://www.umei.cc/p/gaoqing/cn/26.htm")
	fmt.Println(pageStr)
	// 获取连接
	GetImg("https://www.umei.cc/p/gaoqing/cn/26.htm")
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
	// // myTest()
	// url := "http://i1.shaodiyejin.com/small/files/s145.jpg"
	// filename := "1.jpg"
	// Downloads(url, filename)

	// 初始化管道
	chanImageUrls = make(chan string, 100000)
	//  用于监控携程
	chanTask = make(chan string, 26)
	// 2  爬虫携程
	for i := 1; i < 27; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://www.umei.cc/p/gaoqing/cn/" + strconv.Itoa(i) + ".htm")
	}

	// 3 任务统计携程 统计26个任务是否都完成， 完成则关闭管道
	waitGroup.Add(1)
	go checkOk()
	//  4 下载携程 从管道中读取链接并下载
	for i := 0; i < 20; i++ {
		waitGroup.Add(1)
		go DonwloadImg()
	}
	waitGroup.Wait()
}

func GetImg(url string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}

// 下载图片 传入图片叫什么
func Downloads(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get.url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "resp body")
	filename = "./img/" + filename
	//  写出数据
	err = ioutil.WriteFile(filename, bytes, 777)
	if err != nil {
		return false
	} else {
		return true
	}
}

//  并发思路
// 初始化数据管道
// 爬虫写出 26携程想管道中添加图片连接
//  任务统计携程。 检查26个任务是否完成，完成则关闭数据管道
//  下载携程 从管道读取连接下载
var (
	//  存放图片连接的数据管道
	chanImageUrls chan string
	waitGroup     sync.WaitGroup
	chanTask      chan string
)

// 爬图片到管道
//  url是传的整页连接
func getImgUrls(url string) {
	urls := getImgs(url)
	// 遍历切片所有链接 存入数据管道
	for _, url := range urls {
		chanImageUrls <- url
	}

	//  表示当前这个携程完成
	// 每完成一个任务， 写一条数据
	//  用于监控携程已经完成了几个任务
	chanTask <- url
	waitGroup.Done()
}

// 获取当页图片连接
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("公共找到%d条数据\n", len(results))
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

//  任务统计携程
func checkOk() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

// DonwloadImg 下载图片
func DonwloadImg() {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := Downloads(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

// GetFilenameFromUrl 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	//  返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	//  切出来
	filename = url[lastIndex+1:]
	//  时间戳 解决重名
	timePrefix := strconv.Itoa(int(time.Now().Unix()))
	filename = timePrefix + "_" + filename
	return
}
