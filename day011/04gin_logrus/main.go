package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var log = logrus.New()

func initLogrus() (err error) {
	log.Formatter = &logrus.JSONFormatter{} // 设置日志文件

	file, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed  err:", err)
		return
	}
	//  输出日志文件路径
	log.Out = file
	//  告诉gin把它的日志也记录到我们打开的文件中
	gin.SetMode(gin.ReleaseMode) // gin 设置为logrus日志插件
	gin.DisableConsoleColor()
	gin.DefaultWriter = log.Out
	// 设置日志级别
	log.Level = logrus.InfoLevel
	return
}
func main() {
	err := initLogrus()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/index", indexHandler)
	router.Run()
}

func indexHandler(c *gin.Context) {
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Warn("A group of walrus emerges from the ocean")
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"massge": "ok",
	})
}
