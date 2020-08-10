package main

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() (ip string, err error) {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range address {
		ipAddr, ok := addr.(*net.IPNet) // 类型断言
		if !ok {
			continue
		}

		if ipAddr.IP.IsLoopback() {
			continue
		}

		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		fmt.Println(ipAddr)
		return ipAddr.IP.String(), nil
	}
	return
}

// Get preferred outbound ip of this machine
func GetOuboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return strings.Split(localAddr.IP.String(), ":")[0]
}
func main() {
	//GetLocalIP()
	GetOuboundIP()
}
