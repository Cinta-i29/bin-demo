package routes

import (
	"Gin-Demo/controller/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Userinfo 利用context.ShouldBind(&Userinfo)绑定到结构体
// 示例在
type Userinfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"password"`
}

func ApiRoutesInit(router *gin.Engine) {
	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"username": "张三",
				"age":      20})
		})

		apiRoute.GET("/news", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"title": "这是新闻"})
		})

		//绑定到结构体中
		//Get传值/?username=zhangsan&password=123456  Post传值也同理
		apiRoute.GET("/index", func(c *gin.Context) {
			var userinfo Userinfo
			if err := c.ShouldBind(&userinfo); err == nil {
				c.JSON(http.StatusOK, userinfo)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		})

		apiRoute.GET("/add", admin.UserController{}.Add)
	}
}
