package main

import "github.com/gorilla/mux"
import "fmt"
import "net/http"

func main()  {
	// 创建路由 ~/Documents/Learn/goLearn/goProject/src/awesomeProject »
	router := mux.NewRouter()
	// ws控制器不断去处理管道数据，进行同步
	go h.run()
	// 指定ws 的回调函数
	router.HandleFunc("/ws", wsHandler)
	// 开启服务监听
	if err := http.ListenAndServe("127.0.0.1:8080", router); err!= nil{
		fmt.Println("err", err)
	}
}