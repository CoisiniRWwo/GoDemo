package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//创建默认路由引擎
	r := gin.Default()

	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "你好gin")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "我是新闻页面")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "这是一个post请求")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个put请求")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "这是一个DELETE请求")
	})

	//默认在8080端口启动服务
	r.Run(":9001") //启动一个web服务

}
