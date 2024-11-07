package Controller

import (
	"MyTest/Logic/RuleChain/UserRuleChain"
	"MyTest/Logic/log"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Logic/user_managment/UserCreate"
	"MyTest/Models/Error"
	Models "MyTest/Models/Message"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type content struct {
	Content string `json:"content"`
}

// LoadUserMsgPage 加载用户消息界面接口
// @Summary 加载用户消息界面接口
// @Description 加载用户消息界面
// @Tags 加载用户消息界面接口
// @Security ApiKeyAuth
// @Success 200
// @Router /users/rec/{rec_id} [get]
func LoadUserMsgPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		//捕获异常
		defer log.RecoverPanic()
		fmt.Println("load user msg page")
		//获取用户信息
		accountNum, rec_id := c.GetInt64("account_num"), c.Param("rec_id")

		//check操作进入管道
		UserRuleChain.GetUserCheckOpt(accountNum, rec_id)

		//读取要发给前端的数据
		TypeDefine.Mu.Lock()
		Data := <-TypeDefine.UserMap[accountNum].ViewData.Ch
		TypeDefine.Mu.Unlock()

		if Data.Data == nil {
			Error.NewErrHandle(Error.ErrorInit("Data is nil", 400)).WriteErr().ViewErr()
		}
		fmt.Println("get data from channel")

		if Data.Type != "message" {
			Error.NewErrHandle(Error.ErrorInit("Data type error", 400)).WriteErr()
		} else {
			if m, ok := Data.Data.([]Models.Message); ok {
				for _, v := range m {

					c.JSON(http.StatusOK, v.Content)
				}
			}
		}
	}
}

// SendMsgToUser 发送消息给用户接口
// @Summary 发送消息给用户接口
// @Description 发送消息给用户接口
// @Tags 发送消息给用户接口
// @Security ApiKeyAuth
// @Success 200
// @Router /users/rec//{rec_id} [post]
func SendMsgToUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//构造消息体
		Content := content{}
		c.BindJSON(&Content)
		id, err := strconv.Atoi(c.Param("rec_id"))
		if err != nil {
			Error.NewErrHandle(err).WriteErr().ViewErr()
		}
		M := Models.NewMessage(UserCreate.GetFuncMember(c.GetInt64("account_num")).GetID(), uint(id), Content.Content)

		//发送消息给用户
		TypeDefine.Mu.Lock()
		TypeDefine.UserMap[c.GetInt64("account_num")].Opt.Ch <- map[interface{}]interface{}{"opt_type": "sendMsg", "msg": M.Content, "id": M.ReceiverID}
		TypeDefine.Mu.Unlock()
	}
}
