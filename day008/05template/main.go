package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//  template 模板
func info(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./info.html")
	// data, err := ioutil.ReadFile("./info.html")
	if err != nil {
		fmt.Println("open html file failed , err:", err)
		return
	}
	// 用数据去渲染模板
	data := "<li>《我的世界》</li>"
	t.Execute(w, data)
	// num := rand.Intn(10)
	// dataStr := string(data) // 转换成字符串
	// if num > 5 {
	// 	dataStr = strings.Replace(dataStr, "{ooxx}", "<li>《我的世界》</li>", 1)
	// } else {
	// 	dataStr = strings.Replace(dataStr, "{ooxx}", "<li>《对子哈特1枚》</li>", 1)
	// }
	// w.Write([]byte(dataStr))

}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe(":8080", nil)

}
