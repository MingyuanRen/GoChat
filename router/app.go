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
	r.GET("/user/GetUserList", service.GetUserList)
	r.GET("/user/CreateUser", service.CreateUser)
	return r
}
