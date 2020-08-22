package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

// 服务器 求矩形面积和周长
// 声明 矩形对象
type Rect struct {
}

// Rpc 声明参数结构体  字段首字母大写
type Params struct {
	// 长和宽
	Width, Height int
}

// 定义求矩形面积和方法
func (r *Rect) Area(p Params, ret *int) (err error) {
	*ret = p.Width * p.Height
	return nil
}

// 周长方法
func (r *Rect) Perimter(p Params, ret *int) (err error) {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	// 1 注册服务
	rect := new(Rect)
	rpc.Register(rect) // 注册
	// 把服务处理绑定到http协议上面
	rpc.HandleHTTP()
	//	 监听服务 等待客户端调用用于求面积和周长方法
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
