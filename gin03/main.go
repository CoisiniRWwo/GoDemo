package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"score": 98,
			"hobby": []string{"敲代码", "吃饭", "健身"},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻标题111",
					Content: "新闻详情111",
				},
				&Article{
					Title:   "新闻标题222",
					Content: "新闻详情222",
				},
			},
			"testSlice": []string{},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
		})
	})

	router.GET("/news", func(context *gin.Context) {
		news := new(Article)
		news.Title = "新闻标题"
		news.Content = "新闻详情"
		context.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻",
			"news":  news,
		})
	})

	router.GET("/admin", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "首页",
		})
	})

	router.GET("/admin/news", func(context *gin.Context) {
		news := new(Article)
		news.Title = "新闻标题"
		news.Content = "新闻详情"
		context.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "新闻",
			"news":  news,
		})
	})

	router.Run()
}
