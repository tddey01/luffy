package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 上传文件
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./upload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	r.POST("/upload", uploadHandler)
	r.Run()
}

func uploadHandler(c *gin.Context) {
	// 提取用户上传文件 本地存起来
	// 单个文件
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// fileobj, err := c.FormFile("filename")
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"code": 1,
	// 		"msg":  err,
	// 	})
	// 	return
	// }
	// // filaobj 上传的文件对象
	// fmt.Println(fileobj.Filename) // 拿到上传文件的文件名
	// filePath := fmt.Sprintf("./%s", fileobj.Filename)
	// // 保存到本地文件路径
	// c.SaveUploadedFile(fileobj, filePath)
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "ok",
	// })

	//  多文件上传
	fileall, _ := c.MultipartForm()
	files := fileall.File["filename"]
	for index, file := range files {
		fmt.Println(file.Filename)
		dst := fmt.Sprintf("./%s_%d", file.Filename, index)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploaded!", len(files)),
	})
}
