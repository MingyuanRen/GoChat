package router

import (
	"gochat/docs"
	"gochat/service"

	"github.com/gin-gonic/gin"

	// docs "github.com/go-project-name/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/index", service.GetIndex)
	// static
	r.Static("/asset", "asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	r.LoadHTMLGlob("views/**/*")

	// mainpage
	r.GET("/", service.GetIndex)

	// user module
	r.GET("/user/GetUserList", service.GetUserList)
	r.GET("/user/CreateUser", service.CreateUser)
	r.GET("/user/DeleteUser", service.DeleteUser)
	r.POST("/user/UpdateUser", service.UpdateUser)
	r.POST("/user/FindUserByNameAndPwd", service.FindUserByNameAndPwd)

	// send message
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)
	return r
}
