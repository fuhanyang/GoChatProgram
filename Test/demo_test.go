package Test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// handler 处理 /data 请求，返回 JSON 格式的切片数据
func handler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, []string{"apple", "banana", "orange"})
	}
	// 将切片编码为 JSON 并发送响应
}

// Test2 启动一个 HTTP 服务器，并注册 handler
//func Test2(t *testing.T) {
//	defer Mysql.MysqlClose()
//	Mysql.Init()
//	Table.InitTable()
//	router := gin.Default()
//	UsersGroup := router.Group("/users")
//	{
//		UsersGroup.GET(":accountNum", Routers.LoadUsersPage())
//	}
//	//查看别的用户
//	ReceiverGroup := UsersGroup.Group("/rec")
//	{
//		ReceiverGroup.GET("/:rec_id/:accountNum", Routers.LoadUserMsgPage())
//
//	}
//	// 使用 CORS 中间件
//	router.Use(func(c *gin.Context) {
//		c.Header("Access-Control-Allow-Origin", "*")                   // 允许所有域访问
//		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // 允许的 HTTP 方法
//		c.Header("Access-Control-Allow-Headers", "Content-Type")       // 允许的请求头
//
//		// 如果是 OPTIONS 请求，直接返回
//		if c.Request.Method == http.MethodOptions {
//			c.Status(http.StatusOK)
//			return
//		}
//
//		c.Next() // 继续处理请求
//	})
//
//	// 定义 /data 路由
//	router.GET("/data", handler())
//	router.Run(":9090")
//}
