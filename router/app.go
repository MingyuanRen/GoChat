package router

import (
	"gochat/docs"
	"gochat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// static
	r.Static("/asset", "asset/")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//	r.StaticFS()
	r.LoadHTMLGlob("views/**/*")

	// mainpage
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	// r.GET("/chat", service.Chat)
	r.POST("/searchFriends", service.SearchFriends)

	// user module
	r.POST("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.POST("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.POST("/user/find", service.FindByID)
	//send msg
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)
	// add friend
	r.POST("/contact/addfriend", service.AddFriend)
	// upload file
	r.POST("/attach/upload", service.Upload)
	// create group
	r.POST("/contact/createCommunity", service.CreateCommunity)
	// group list
	r.POST("/contact/loadcommunity", service.LoadCommunity)
	r.POST("/contact/joinGroup", service.JoinGroups)

	r.POST("/user/redisMsg", service.RedisMsg)
	return r
}
