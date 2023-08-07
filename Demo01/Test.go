package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitAdminMiddleware(ctx *gin.Context) {
	fmt.Println("路由分组中间件")
	// 可以通过 ctx.Set 在请求上下文中设置值，后续的处理函数能够取到该值ctx.Set("username", "张三")
	// 调用该请求的剩余处理程序
	ctx.Next()
}
