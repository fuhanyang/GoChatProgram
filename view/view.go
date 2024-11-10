package view

import (
	"MyTest/Logic/message_systerm"
	"MyTest/Logic/user_managment/TypeDefine"
	Models "MyTest/Models/Message"
	"MyTest/Models/Users/FunctionalMember"
)

func ShowMsgReceiver(...interface{}) {

}

// PrintMsg 数据通过管道给前端
func PrintMsg(m []Models.Message, account int64) {
	TypeDefine.UserMap[account].ViewData.Ch <- TypeDefine.ViewData{message_systerm.Msg, m}
}
func PrintUser([]FunctionalMember.FuncMember) {

}
