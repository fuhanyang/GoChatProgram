package Controller

import (
	"MyTest/Logic/Notice"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Logic/user_managment/UserHandle"
	"MyTest/Models/Error"
	"MyTest/view"
	"github.com/gin-gonic/gin"
	"strconv"
)

// LoadUsersPage 加载用户信息界面接口
// @Summary 加载用户信息界面接口
// @Description 加载用户信息界面
// @Tags 加载用户信息界面接口
// @Security ApiKeyAuth
// @Success 200
// @Router /users// [get]
func LoadUsersPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		//捕获异常
		defer Notice.RecoverPanic()

		accountStr := c.Param(TypeDefine.AccountNum)
		accountNum, err := strconv.ParseInt(accountStr, 10, 64)
		Error.NewErrHandle(err).WriteErr()

		//显示用户信息
		view.ShowMsgReceiver(UserHandle.GetMsgReceiver(accountNum))

	}
}
