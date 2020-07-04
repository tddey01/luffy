package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// User 用户
type User struct {
	UserName string
	Password string
	Age      int
}

//  template 模板
func info(w http.ResponseWriter, r *http.Request) {
	// 添加自定义的方法要在parse模板文件之前添加

	// 1. 自定义一个函数
	// 自定义一个夸人的模板函数
	kuaFunc := func(arg string) (string, error) {
		return arg + "真帅", nil
	}

	
	t, err := template.ParseFiles("./info1.html")
	// data, err := ioutil.ReadFile("./info.html")
	if err != nil {
		fmt.Println("open html file failed , err:", err)
		return
	}

	// 用数据去渲染模板
	// data := "<li>《我的世界》</li>"
	// t.Execute(w, data)

	// 	user := User{
	// 		"北京",
	// 		"1234",
	// 		19,
	// 	}
	// 	t.Execute(w, user)
	userMap := map[int]User{
		// userMap := []User{
		1: {"北京", "1234", 18},
		2: {"上海", "1234", 38},
		3: {"天津", "1234", 28},
	}
	t.Execute(w, userMap)
}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe(":8080", nil)

}
