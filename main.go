package main

import (
	"Gin-Demo/models"
	"Gin-Demo/routes"
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
)

// 定义全局处理器
func initMiddleware(ctx *gin.Context) {
	//进行业务处理 例如身份验证、日志记录等等
	fmt.Println("全局中间件 通过 r.Use 配置")
	//调用该请求的剩余处理程序
	ctx.Next()
}

func main() {
	//分配一个默认的路由
	r := gin.Default()

	//md5加密演示
	h := md5.New()
	_, err2 := io.WriteString(h, "123456")
	if err2 != nil {
		return
	}
	fmt.Printf("%x\n\n", h.Sum(nil))

	//=sessions
	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret11111"))
	// 设置 session 中间件，参数 MySession，指的是 session 的名字，也是cookie 的名字
	// store 是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("MySession", store))

	//注册全局模板函数
	r.SetFuncMap(template.FuncMap{"unixToDate": models.UnixToDate})
	//应用全局处理器
	r.Use(initMiddleware)
	//加载模板 渲染html
	r.LoadHTMLGlob("templates/**/*")
	//实现路由分组
	routes.AdminRoutesInit(r)
	routes.UserRoutesInit(r)
	routes.ApiRoutesInit(r)
	routes.DefaultRoutesInit(r)
	_ = r.Run(":8080")
}
