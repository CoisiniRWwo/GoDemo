package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {

	r := gin.Default()
	//配置模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(x *gin.Context) {
		x.String(http.StatusOK, "值：%v", "首页")
	})

	r.GET("/json1", func(x *gin.Context) {
		x.JSON(http.StatusOK, map[string]interface{}{
			"success": "true",
			"msg":     "你好，gin1",
		})
	})

	r.GET("/json2", func(x *gin.Context) {
		x.JSON(http.StatusOK, gin.H{
			"success": "true",
			"msg":     "你好，gin2",
		})
	})

	r.GET("/json3", func(x *gin.Context) {
		a := &Article{
			Title:   "我是帅哥",
			Desc:    "描述",
			Content: "测试内容",
		}
		x.JSON(http.StatusOK, a)
	})

	//响应JSONP请求
	//	http://localhost:9001/json4?callback=xxx
	r.GET("/json4", func(x *gin.Context) {
		a := &Article{
			Title:   "我是帅哥",
			Desc:    "描述",
			Content: "测试内容",
		}
		x.JSONP(http.StatusOK, a)
	})

	r.GET("/xml", func(x *gin.Context) {
		x.XML(http.StatusOK, gin.H{
			"success": "true",
			"mag":     "你好，这是xml",
		})
	})

	r.GET("/news", func(x *gin.Context) {
		//需要r.LoadHTMLGlob("templates/*")
		x.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})

	r.GET("/index", func(x *gin.Context) {
		//需要r.LoadHTMLGlob("templates/*")
		x.HTML(http.StatusOK, "index.html", gin.H{
			"title": "我是后台数据",
		})
	})

	r.Run(":9001")

}
