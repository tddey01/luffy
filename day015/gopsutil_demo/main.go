package main

import (
	"fmt"
	"log"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/shirou/gopsutil/cpu"
)

// influxdb demo
var (
	cli client.Client
)

func initconnInflux() (err error) {
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://172.16.0.104:8086",
		Username: "admin",
		Password: "",
	})
	return
}

// insert
func writesPoints(percent int64) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	tags := map[string]string{"cpu": "cpu0"}
	fields := map[string]interface{}{
		"cpu_percent": percent,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert success")
}

func getcpu() {
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
	writesPoints(int64(percent[0]))
}

func main() {
	err := initconnInflux()
	if err != nil {
		fmt.Printf("conect to influxdb failed err :%v", err)
		return
	}
	//	 每一秒钟执行一次
	for {
		getcpu()
		time.Tick(time.Second)
	}
}
