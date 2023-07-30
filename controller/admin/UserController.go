package admin

import (
	"Gin-Demo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type UserController struct {
}

func (c UserController) Index(ctx *gin.Context) {
	//得到共享数据
	name, exists := ctx.Get("name")
	if exists {
		fmt.Println("共享数据=>得到name:", name)
	}
	ctx.String(http.StatusOK, "这是用户首页")
}

func (c UserController) DoAdd(ctx *gin.Context) {
	username := ctx.PostForm("username")
	//单文件
	file, err := ctx.FormFile("face")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	/**
	//多文件下采用下列方式
		//也是在html中加入enctype="multipart/form-data"标签
		form, _ := ctx.MultipartForm()
		files := form.File["face[]"]
		for _, file := range files {
			// 上传文件至指定目录
			dst := path.Join("./static/upload", file.Filename)
			ctx.SaveUploadedFile(file, dst)}
		}
	*/
	// 上传文件到指定的目录
	dst := path.Join("./static/upload", file.Filename)
	fmt.Println(dst)
	err2 := ctx.SaveUploadedFile(file, dst)
	if err2 != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("'%s' uploaded!", file.Filename), "username": username})
}

func (c UserController) Add(ctx *gin.Context) {
	////这里写死了
	//user := models.User{
	//	Username: "itying.com",
	//	Age:      18,
	//	Email:    "itying@qq.com",
	//	AddTime:  int(time.Now().Unix()),
	//}
	//result := models.DB.Create(&user) // 通过数据的指针来创建
	//if result.RowsAffected > 1 {
	//	fmt.Print(user.Id)
	//}
	//fmt.Println(result.RowsAffected)
	//fmt.Println(user.Id)
	//ctx.String(http.StatusOK, "add 成功")
	//
	var users []models.User
	models.DB.Where("id<10").Find(&users)
	ctx.JSON(http.StatusOK, gin.H{
		"result": users,
	})
}

// SelectId 查询用户信息
func (c UserController) SelectId(ctx *gin.Context) {
	id, exists := ctx.Get("uid")
	if !exists {
		println("===========================================")
		println("没有获取到uid")
		println("===========================================")
	}
	user := models.User{}
	models.DB.Where("id =? ", id).Find(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (c UserController) AddUser(ctx *gin.Context) {
	userinfo, _ := ctx.Get("adduser")
	models.DB.Create(userinfo)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户添加成功",
	})
}

func (c UserController) UpdateUser(ctx *gin.Context) {
	update, _ := ctx.Get("updateuser")
	models.DB.Updates(update)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户修改成功",
	})
}

func (c UserController) Delete(ctx *gin.Context) {
	id, _ := ctx.Get("deleteid")
	result := models.DB.Where("id = ?", id).Delete(&models.User{})
	// 检查删除操作是否出现错误
	if err := result.Error; err != nil {
		// 处理删除操作出现的错误
		// 例如，可以向客户端返回错误信息或者记录错误日志
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除用户时发生了错误",
		})
		return
	}

	// 检查是否有记录被删除
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		// 处理没有找到记录的情况
		// 例如，可以向客户端返回一个消息，表示没有找到要删除的用户
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "未找到要删除的用户",
		})
		return
	}

	// 在这里可以处理删除成功的情况，例如向客户端返回一个成功的响应
	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户删除成功",
	})
}
