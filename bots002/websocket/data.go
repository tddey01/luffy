package main

// 将连接中传输的数据抽象出对象

// Data  结构体
type Data struct {
	Ip       string   `json:"ip"`
	Type     string   `json:"type"`      // 表示信息类型  login 登录 handshake 握手信息 刚打开网页的状态  system 系统信息 user 普通信息   Logout  退出信息
	From     string   `json:"from"`      // 代表那个用户说的
	Content  string   `json"content"`    // 传输内容
	User     string   `json:"user"`      // 用户名
	UserList []string `json:"user_list"` // 用户列表
}
