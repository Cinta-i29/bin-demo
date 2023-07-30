package routes

import (
	"Gin-Demo/controller/admin"
	"Gin-Demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutesInit(router *gin.Engine) {
	adminRouter := router.Group("/admin")
	{
		//添加用户的页面
		adminRouter.GET("/user/add", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin/user/add.html", gin.H{})
		})

		//添加用户请求 这里主要演示如何添加文件
		adminRouter.POST("/user/doAdd", admin.UserController{}.DoAdd)

		//调用Controller
		adminRouter.GET("/news", admin.NewsController{}.Index)

		//返回http
		adminRouter.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin/index.html", gin.H{
				//这里使用到了全局模板函数
				"now":   models.GetUnix(),
				"title": "admin",
			})
		})
	}
}
