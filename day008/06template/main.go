package main

import (
	"fmt"
	"io/ioutil"
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
	// 打开一个模板文件
	htmlByte, err := ioutil.ReadFile("./info1.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 添加自定义的方法要在parse模板文件之前添加

	// 1. 自定义一个函数
	// 自定义一个夸人的模板函数
	kuaFunc := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 2 把自定义的函数告诉模板系统
	// template.New("info") // 创建一个Template对象
	// template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}) // 给模板系统追加自定义函数

	// 解析模板

	// 链式操作？
	// 原理：每一次执行完方法之后返回操作的对象本身
	t, err := template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(htmlByte))
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
