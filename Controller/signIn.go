package Controller

import (
	"MyTest/Logic/Notice"
	"MyTest/Logic/user_managment/UserCreate"
	"MyTest/Logic/user_managment/UserHandle"
	"MyTest/Models/Error"
	"MyTest/Models/Users/FunctionalMember"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoadEntryPage 加载用户登录界面接口
// @Summary 加载用户登录界面接口
// @Description 加载用户登录界面
// @Tags 加载用户登录界面接口
// @Security ApiKeyAuth
// @Success 200
// @Router /pri/entry/ [get]
func LoadEntryPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
	}
}

// UserSignIn 用户登录接口
// @Summary 用户登录接口
// @Description 用户登录接口
// @Tags 用户登录接口
// @Security ApiKeyAuth
// @Param user body FunctionalMember.FuncMember true "用户登录信息"
// @Success 200
// @Router /pri/entry/ [post]
func UserSignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		//捕获异常
		defer Notice.RecoverPanic()
		//用户登录，同步数据
		var UserData FunctionalMember.FuncMember
		if err := c.BindJSON(&UserData); err != nil {
			Error.NewErrHandle(err).WriteErr()
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  "无效的参数",
			})
			return
		}
		M, err := UserCreate.UserSignIn(UserData.PassWord, UserData.AccountNum, c.ClientIP())
		Error.NewErrHandle(err).WriteErr().ViewErr()

		if M != nil {
			//用户上线
			Error.NewErrHandle(UserHandle.UserOnline(M)).WriteErr().ViewErr()
			//显示用户信息
		} else {
			c.JSON(http.StatusConflict, gin.H{"password": "wrong"})
		}

		//jwt验证
		// 校验用户名和密码是否正确
		M = UserCreate.GetFuncMember(UserData.AccountNum)
		if M == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  "用户不存在",
			})
		}
		if UserData.AccountNum == M.AccountNum && UserData.PassWord == M.PassWord {
			// 生成Token
			tokenString, _ := GenToken(UserData.Name, UserData.AccountNum)
			c.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "success",
				"data": gin.H{"token": tokenString},
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2002,
			"msg":  "鉴权失败",
		})
		return
	}
}
