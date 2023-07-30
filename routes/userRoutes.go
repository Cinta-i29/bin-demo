package routes

import (
	"Gin-Demo/controller/admin"
	"Gin-Demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserInfo struct {
	Id       int    `from:"id" json:"id"`
	Username string `from:"username" json:"username"`
	Age      int    `from:"age" json:"age"`
	Email    string `from:"email" json:"email"`
	AddTime  int    `from:"addtime" json:"addtime"`
}

func UserRoutesInit(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		// user/1
		userRouter.GET("/:uid", func(ctx *gin.Context) {
			uid := ctx.Param("uid")
			ctx.Set("uid", uid)
		}, admin.UserController{}.SelectId)

		userRouter.POST("/add", func(ctx *gin.Context) {
			var user models.User
			if err := ctx.ShouldBind(&user); err == nil {
				ctx.Set("adduser", &user)
				admin.UserController{}.AddUser(ctx)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		})

		userRouter.PUT("/update", func(ctx *gin.Context) {
			var user models.User
			if err := ctx.ShouldBind(&user); err == nil {
				ctx.Set("updateuser", &user)
				admin.UserController{}.UpdateUser(ctx)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		})

		userRouter.DELETE("/delete/:uid", func(ctx *gin.Context) {
			uid := ctx.Param("uid")
			id, err := strconv.Atoi(uid)
			if err != nil {
				return
			}
			user := models.User{Id: id}
			ctx.Set("deleteid", user.Id)
		}, admin.UserController{}.Delete)
	}
}
