package Controller

import (
	"MyTest/Logic/Notice"
	"MyTest/Logic/user_managment/UserCreate"
	"MyTest/Models/Error"
	"MyTest/Models/Users/User"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoadRegiPage 加载用户注册界面接口
// @Summary 加载用户注册界面接口
// @Description 加载用户注册界面
// @Tags 加载用户注册界面接口
// @Security ApiKeyAuth
// @Success 200
// @Router /pri/regi/ [get]
func LoadRegiPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "signup page load success",
		})
	}
}

// UserRegister 用户注册接口
// @Summary 用户注册接口
// @Description 用户注册接口
// @Tags 用户注册接口
// @Security ApiKeyAuth
// @Param user body User.UserInfo true "用户注册信息"
// @Success 200
// @Router /pri/regi/ [post]
func UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		//捕获异常
		defer Notice.RecoverPanic()

		ip := c.ClientIP()
		fmt.Println(ip)
		var UserInfo User.UserInfo
		Error.NewErrHandle(c.BindJSON(&UserInfo)).WriteErr().ViewErr()
		fmt.Println(UserInfo)
		M := UserCreate.UserSignUp(ip, UserInfo.PassWord, UserInfo.Name)
		if M == nil {
			Error.NewErrHandle(Error.ErrorInit("User signs up error", 400)).WriteErr().ViewErr()
		}

	}
}
