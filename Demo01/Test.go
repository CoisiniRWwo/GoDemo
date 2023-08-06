package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
func main() {
	router := gin.Default()
	router.GET("/addUser", func(c *gin.Context) {
		c.HTML(200, "default/add_user.html", gin.H{})
	})
	router.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		c.JSON(200, gin.H{"usernmae": username, "password": password, "age": age})
	})
}
