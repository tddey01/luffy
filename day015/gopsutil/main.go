package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/host"

	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/load"

	"github.com/shirou/gopsutil/cpu"
)

func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed ,err:%v\n", err)
		return
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

// cpu 负载
func getLoad() {
	info, err := load.Avg()
	if err != nil {
		fmt.Printf("load.Avg() failed err:%v\n", err)
		return
	}
	fmt.Println(info)
}

// 内存
func getMeminfo() {
	meminfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", meminfo)
}

// host
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}
func main() {
	//getCpuInfo()
	//getLoad()
	//getMeminfo()
	getHostInfo()
}
