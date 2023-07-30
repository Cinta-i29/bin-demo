package admin

import (
	"Gin-Demo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsController struct {
	//控制器继承
	BaseController
}

func (c NewsController) Index(ctx *gin.Context) {
	//调用了父类控制器的Success方法
	c.Success(ctx)
	//调用了模板 也就是工具类
	fmt.Println(models.GetDay())
	ctx.String(http.StatusOK, "新闻首页")
}
