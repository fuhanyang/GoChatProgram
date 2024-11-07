package Routers

import (
	"MyTest/Controller"
	_ "MyTest/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

var router *gin.Engine

func RouterDefault() *gin.Engine {

	router = gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLFiles("")
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	return router
}

func PriRouters() {

	PriGroup := router.Group("/pri")
	PriGroup.GET("/", Controller.LoadPriPage())

	RegiGroup := PriGroup.Group("/regi")
	{
		RegiGroup.GET("/", Controller.LoadRegiPage())
		RegiGroup.POST("/", Controller.UserRegister())
	}
	EntryGroup := PriGroup.Group("/entry")
	{
		EntryGroup.GET("/", Controller.LoadEntryPage())
		EntryGroup.POST("/", Controller.UserLoginIn())
	}
	GuestGroup := PriGroup.Group("/guest")
	{
		GuestGroup.GET("/")
		GuestGroup.POST("/")

	}
}

// 用户
func UserRouters() {
	UsersGroup := router.Group("/users")

	//jwt 中间件
	UsersGroup.Use(Controller.JWTAuthMiddleware())

	{
		UsersGroup.GET(":accountNum", Controller.LoadUsersPage())
	}
	//查看别的用户的信息的一些接口
	ReceiverGroup := UsersGroup.Group("/rec")

	//jwt 中间件
	ReceiverGroup.Use(Controller.JWTAuthMiddleware())

	{
		//查看和别的用户的对话信息
		ReceiverGroup.GET("/:rec_id", Controller.LoadUserMsgPage())
		//发消息给别的用户
		ReceiverGroup.POST("/:rec_id", Controller.SendMsgToUser())
	}
}

func Init() *gin.Engine {
	router = RouterDefault()
	UserRouters()
	PriRouters()
	return router
}
