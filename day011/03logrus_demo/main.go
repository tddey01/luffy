package main

import "github.com/sirupsen/logrus"

// logrus示例
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{}) // 设置json格式日志
	logrus.SetReportCaller(true)
	logrus.WithFields(logrus.Fields{
		"name": "官大码",
		"age":  9000,
	}).Warn("这是一条warning级别日志")
	logrus.Info("这是一条INFO日志")
}
