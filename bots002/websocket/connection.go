package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// 需要的数据结构
// ws连接器

type connection struct {
	// ws 连接器
	ws *websocket.Conn
	//	 管道
	send chan []byte
	//	 数据
	data Data
}

// ws连接器
// 处理ws中的各种逻辑
type hub struct {
	//	connection 注册连接器
	connection map[*connection]bool
	//	从连接器发送的信息
	broadcast chan []byte
	//	从连机器注册请求
	reqister chan *connection
	//	从连机器销毁请求
	unregister chan *connection
}

// ws的读和写
// 1 往ws中写数据
func (c *connection)writeToWs()  {
	// 从管道遍历数据
	for message := range c.send{
		// 数据写出
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var user_list = []string{}

// ws 连接中读数据
func (c *connection) reader() {
	//  不断读取数据
	for {
		_, massge, err := c.ws.ReadMessage()
		if err != nil {
			// 读不进数据，将这个用户移除队列
			h.unregister <- c
			break
		}
		//	 读取数据
		json.Unmarshal(massge, c.data)
		// 跟进data中的type判断应该进行什么操作
		switch c.data.Type {
		case "login":
			// 弹出窗口，输入用户名
			c.data.User = c.data.Content
			c.data.From = c.data.User
			// 登录后将用户加入用户列表
			user_list = append(user_list, c.data.User)
			// 每一个登录的用户都要看到所有已经登录的用户
			c.data.UserList = user_list
			// 数据序列化
			data_b, err := json.Marshal(c.data)
			if err != nil {
				fmt.Println(err)
			}
			h.broadcast <- data_b
			// 普通用户
		case "user":
			c.data.Type = "user"
			data_b, err := json.Marshal(c.data)
			if err != nil {
				fmt.Println(err)
			}
			h.broadcast <- data_b
		case "logot":
			c.data.Type = "user"
			// 用户列表删除
			user_list = removeUser(user_list, c.data.User)
			c.data.UserList = user_list
			c.data.Content = c.data.User
			// 数据序列化，让所有人知道xx下线了
			data_b, err := json.Marshal(c.data)
			if err != nil {
				fmt.Println(err)
			}
			h.broadcast <- data_b
			h.unregister <- c
		default:
			fmt.Println("其他")
		}
	}
}

// 删除用户 删除用户切片中的数据
func removeUser(slice []string, user string) []string {
	// 严谨判断
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	// 定义新的返回切片
	var my_slice = []string{}
	// 删除传入切片中的指定用户，其他的用户放到新的切片中
	for i := range slice {
		// 利用索引删除用户
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			my_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return my_slice
}

// 定义一个升级器 将http 升级为ws请求
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ws的回调函数
// http.ResponseWriter 响应
// *http.Request 请求
func wsHandler(w http.ResponseWriter, r *http.Request)  {
	ws,err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		fmt.Println(err)
	}
	// 创建链接对象 去做事情
	// 初始化连接对象
	c := &connection{send: make(chan []byte, 128), ws:ws, data: &Data{}}
	// 在ws中注册一下
	h.reqister <- c
	// ws 将数据读写跑起来
	go c.writeToWs()
	c.reader()

	//当主函数执行完毕，执行
	defer func() {
		c.data.Type = "logout"
		// 用户列表删除
		user_list = removeUser(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		// 数据序列化，让所有人知道xx下线了
		data_b, err := json.Marshal(c.data)
		if err != nil{
			fmt.Println(err)
		}
		h.broadcast <- data_b
		h.unregister <- c
		//
	}()

}