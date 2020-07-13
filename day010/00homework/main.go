package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// 登录作业
// 注册
//   实现用户注册用户名和密码 并写入数据库
// 登录
// 	实现用户名和密码登录， 去数据库验证用户和密码是和否正确

// 注册函数
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// 根绝请求方法不同， 来做不同的chuli
	// 如果是POST请求， 就是提取用户提交的form表单数据， 去数据库创建一行数据

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		password := r.FormValue("'password'")
		// 网数据库里面写入数据
		err = createUser(username, password)
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
		}

	} else {
		// 如果是GET请求， 就返回一个HTML页面， 功用户输入注册信息
		t, err := template.ParseFiles("./register.html")
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
		}
		t.Execute(w, nil)
	}

}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/register", registerHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动http server 失败", err)
	}
}
