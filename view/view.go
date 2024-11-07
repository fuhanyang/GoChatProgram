package view

import (
	"MyTest/Logic/user_managment/TypeDefine"
	Models "MyTest/Models/Message"
	"MyTest/Models/Users/FunctionalMember"
	"fmt"
)

func ShowMsgReceiver(...interface{}) {

}

// PrintMsg 数据通过管道给前端
func PrintMsg(m []Models.Message, account int64) {
	for _, v := range m {
		fmt.Println(v.Content)
	}
	TypeDefine.UserMap[account].ViewData.Ch <- TypeDefine.ViewData{"message", m}
}
func PrintUser([]FunctionalMember.FuncMember) {

}
