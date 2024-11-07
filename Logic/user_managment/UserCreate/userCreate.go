package UserCreate

import (
	"MyTest/DAO/Mysql"
	"MyTest/Logic/log"
	"MyTest/Logic/user_managment/Snowflake"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Logic/user_managment/UserHandle"
	"MyTest/Models/Error"
	"MyTest/Models/Users/FunctionalMember"
	"MyTest/rabbitmq"
	"fmt"
)

// GetFuncMember 找到账号对应的用户对象
func GetFuncMember(AccountNum int64) *FunctionalMember.FuncMember {
	//
	var M FunctionalMember.FuncMember

	Mysql.MysqlDb.Model(&FunctionalMember.FuncMember{}).First(&M, FunctionalMember.FuncMember{AccountNum: AccountNum})

	return &M
}

// 用户的注册登录

// CreateAccountNum 产生账号
func CreateAccountNum() int64 {
	//捕获异常
	defer log.RecoverPanic()
	AccountNum := Snowflake.GetID()
	fmt.Println("account num is ", AccountNum)
	return AccountNum
}

// UserSignUp 注册
func UserSignUp(IP string, PassWord string, Name string) (M *FunctionalMember.FuncMember) {
	//捕获异常
	defer log.RecoverPanic()

	M = FunctionalMember.NewFuncMember(Name, IP, CreateAccountNum(), PassWord)
	if M != nil {
		Error.NewErrHandle(M.UserSignUp()).WriteErr().ViewErr()
		//mq
		rabbitmq.UserMsgQueueDeclare(M.ID, M.Name)
	} else {
		Error.NewErrHandle(Error.ErrorInit("User sign up error!", 400)).WriteErr().ViewErr()
	}
	fmt.Println("user sign up " + IP + PassWord)
	return M
}

// UserSignIn 登录
func UserSignIn(PassWord string, AccountNum int64, ip string) (*FunctionalMember.FuncMember, error) {
	//捕获异常
	defer log.RecoverPanic()

	M := GetFuncMember(AccountNum)
	if M == nil {
		return nil, Error.ErrorInit("M is nil!", 400)
	}
	if PassWord == M.PassWord {
		Error.NewErrHandle(M.UserSignIn()).WriteErr().ViewErr()
		M.ChangeIp(ip)

		//获取用户操作
		obj := TypeDefine.UserObj{
			Opt:          UserHandle.GetOptCh(*M),
			IdentityCode: UserHandle.UserIdentityCode(UserHandle.GetAccountNum(M.GetID()), M.GetIP()),
			ViewData:     UserHandle.GetViewDataCh(*M),
		}

		TypeDefine.Mu.Lock()
		TypeDefine.UserMap[UserHandle.GetAccountNum(M.GetID())] = obj
		TypeDefine.Mu.Unlock()

		//日志写入
		log.WriteLog("FuncMember sign in ", fmt.Sprintf("accountNum:%d Name:%s", M.AccountNum, M.Name))

		return M, nil
	} else {
		fmt.Println(M.PassWord)
		return nil, Error.ErrorInit("Wrong password", 400)
	}
}

// UserLogOff 下线
func UserLogOff(PassWord string, AccountNum int64) {
	M := GetFuncMember(AccountNum)
	if PassWord == M.PassWord {
		M.UserLogOff()

	}
}
