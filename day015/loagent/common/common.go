package common

import (
	"fmt"
	"net"
	"strings"
)

const (
	ConNotGetIP = "get.ip.failed"
)

// CollectEntry 要收集的日志配置项结构体
type CollectEntry struct {
	Path  string `json:"path"`  // 去哪个路径去读日志文件
	Topic string `json:"topic"` // 日志文件发往kafka中的那个topic
}

// Get preferred outbound ip of this machine
func GetOuboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
