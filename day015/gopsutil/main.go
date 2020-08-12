package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/net"

	"github.com/shirou/gopsutil/disk"

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

// 磁盘
func getdisk() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	//for _, part := range parts {
	//	fmt.Printf("part:%v\n", part.String())
	//	diskInfo, _ := disk.Usage(part.Mountpoint)
	//	fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	//}
	//
	//ioStat, _ := disk.IOCounters()
	//for k, v := range ioStat {
	//	fmt.Printf("%v:%v\n", k, v)
	//}
	fmt.Println(parts)
	for _, part := range parts {
		partInfo, err := disk.Usage(part.Mountpoint)
		if err != nil {
			fmt.Printf("get part stat failed,err:%v\n", err)
			return
		}
		fmt.Println(partInfo)
	}
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

// 网络net IO
func getnetworkinfo() {
	info, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net io counters failed , err:%v\n", err)
		return
	}
	fmt.Println(info)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

func main() {
	//getCpuInfo()
	//getLoad()
	//getMeminfo()
	//getHostInfo()
	//getdisk()
	getnetworkinfo()
}
