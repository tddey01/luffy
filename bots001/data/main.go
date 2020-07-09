package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"strings"
)

//  数据清洗

// // 读数据
// func read() {

// 	contentBytes, err := ioutil.ReadFile("./data.txt")
// 	if err != nil {
// 		fmt.Println("读入失败", err)
// 	}
// 	contentStr := string(contentBytes)
// 	//  逐行打印 并处理乱码
// 	lineStrs := strings.Split(contentStr, "\n\r") // 切片
// 	for _, lineStr := range lineStrs {
// 		newStr := convertEncoding(lineStr, "GBK")
// 		fmt.Println(newStr)
// 	}

// }

// // 方法2  缓冲读取
// func read2() {
// 	file, err := os.Open("./data.txt")
// 	if err != nil {
// 		fmt.Println("打开文件失败", err)
// 	}
// 	defer file.Close()
// 	//  先建立缓冲区
// 	reader := bufio.NewReader(file)
// 	for {
// 		linebytes, _, err := reader.ReadRune()
// 		if err == io.EOF {
// 			break
// 		}
// 		gbkStr := string(linebytes)
// 		utfStr := convertEncoding(gbkStr, "GBK")
// 		fmt.Println(utfStr)
// 	}

// }

// // 解决字符编码
// // 参数 1 处理的数据
// // 参数 2 数据目前编码
// // 参数 3 返回的正常数据
func convertEncoding(srcStr string, encoding string) (dstStr string) {
	// 创建编码处理器
	enc := mahonia.NewEncoder(encoding)
	// 编码器处理字符串为utf8的字符串
	utfStr := enc.ConvertString(srcStr)
	dstStr = utfStr
	return dstStr
}

func main() {
	// read2()
	file, _ := os.Open("./date.txt")
	// 创建优质文件
	goodFile ,_:= os.OpenFile("./data1.txt", os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0666)
	defer goodFile.Close()

	// 创建劣质文件
	bodFile ,_:= os.OpenFile("./data2.txt", os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0666)
	defer bodFile.Close()

	// 缓冲读取
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		gbkStr := string(lineBytes)
		lineStr := convertEncoding(gbkStr, "GBK")
		//  根据行数， 取身份证
		filelds := strings.Split(lineStr, ",")
		//  判断长度大于等于2 下标1的位置长度=18
		if len(filelds) >= 2 && len(filelds[1])==18{
			goodFile.WriteString(lineStr + "\n")
			fmt.Println("Good", lineStr)
		} else {
			bodFile.WriteString(lineStr + "\n")
			fmt.Println("bod", lineStr)
		}

	}
}
