package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//  HTTP Server 嗨
func main() {
	http.HandleFunc("/web", search)
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./form.html")
	if err != nil {
		fmt.Println("read html file failed , err", err)
		return
	}
	w.Write(data)
}

func index(w http.ResponseWriter, r *http.Request) {
	//  获取注册的信息
	//  获取请求的方法
	//  r:代表跟请求相关的所有内容
	fmt.Println(r.Method)
	r.ParseForm() // 解析
	//  获取表单中的数据
	fmt.Printf("%#v\n", r.Form)
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("pwd"))

	w.Write([]byte("index"))
}
