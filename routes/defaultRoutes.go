package routes

import (
	"Gin-Demo/controller/admin"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 路由中间件
// 钩子函数Hook
func initMiddleware(ctx *gin.Context) {
	//实现共享数据
	ctx.Set("name", "zhangsan")

	fmt.Println("1-执行中间件")
	start := time.Now().UnixNano()
	//.Abort() 终止调用该请求的剩余处理程序
	//.Next() 调用该请求剩余的处理程序
	ctx.Next()
	fmt.Println("3-程序执行完成，计算时间")
	end := time.Now().UnixNano()
	fmt.Println("执行所花费的时间是", end-start)
}

func DefaultRoutesInit(router *gin.Engine) {
	//为整个路由组设置路由中间件
	//写法1：defaultRoute := router.Group("/", initMiddleware)
	//写法2：defaultRoute.Use(initMiddleware)
	defaultRoute := router.Group("/")
	{
		//为一个路由设置路由中间件
		defaultRoute.GET("/", initMiddleware, func(c *gin.Context) {
			fmt.Println("2-执行首页返回数据")
			c.String(200, "首页")
		})
		//initMiddleware和Index实现数据共享
		defaultRoute.GET("/index", initMiddleware, admin.UserController{}.Index)

	}
}
