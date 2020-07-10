package main

import "encoding/json"

// 将连接器对象初始化
var h = hub{
	//	connection 注册连接器
	connection: make(map[*connection]bool),
	//	从连接器发送的信息
	broadcast: make(chan []byte),
	//	从连机器注册请求
	reqister: make(chan *connection),
	//	从连机器销毁请求
	unregister: make(chan *connection),
}

// 处理ws的逻辑实现
func (h *hub) run()  {
	// 监听数据管道，正在后端处理管道数据
	for {
		//跟进不同的数据管道，处理不同的逻辑
		select {
		// 注册
		case c:= <- h.reqister:
			// 标注注册
			h.connection[c] = true
			// 组装data数据
			c.data.Ip = c.ws.RemoteAddr().String()
			// 更改类型
			c.data.Type = "handshake"
			// 更新用户列表
			c.data.UserList = user_list
			data_b,_ := json.Marshal(c.data)
			// 将数据放入管道
			c.send <- data_b
		case c := <- h.unregister:
			// 注销
			if _,ok := h.connection[c]; ok{
				delete(h.connection, c)
				close(c.send)
			}
		case data := <- h.broadcast:
			// 处理数据流转，将数据同步到所有的用户,遍历所有的连接
			// c是单个连接
			for c:= range h.connection{
				// 将数据同步
				select {
				case c.send <- data:
				default:
					// 防止死循环
					delete(h.connection, c)
					close(c.send)
				}
			}
		default:

		}
	}
}