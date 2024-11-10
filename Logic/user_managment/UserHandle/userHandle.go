package UserHandle

import (
	"MyTest/DAO/Mysql"
	"MyTest/Define"
	"MyTest/Logic/Notice"
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/RuleChain/Default"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	"MyTest/Models/Log"
	"MyTest/Models/Users/FunctionalMember"
	"MyTest/Models/Users/User"
	"context"
	"fmt"
	"strconv"
	"time"
)

// GetAccountNum 从数据库查找id对应的账号
func GetAccountNum(id uint) int64 {
	//
	var M FunctionalMember.FuncMember
	Mysql.MysqlDb.First(&M, id)
	return M.AccountNum
}

// UserIdentityCode 识别码
func UserIdentityCode(accountNum int64, IP string) string {
	accountStr := strconv.FormatInt(accountNum, 10)

	return accountStr + IP
}
func GetUserParams(u User.UserInf) (params RuleChain.Params) {
	user, ok := u.(*FunctionalMember.FuncMember)
	params = make(RuleChain.Params)
	if !ok {
		//游客
	} else {

		params[TypeDefine.Ch] = TypeDefine.UserMap[GetAccountNum(user.GetID())].Opt.Ch
		params[TypeDefine.Name] = user.Name
		params[TypeDefine.AccountNum] = user.AccountNum
		params[TypeDefine.Password] = user.PassWord
		params[TypeDefine.IP] = user.IP
		params[TypeDefine.ID] = user.ID

	}
	return params
}

// GetOptCh 读取用户操作的管道
func GetOptCh(u FunctionalMember.FuncMember) TypeDefine.UserOpt {
	ch := make(TypeDefine.UserOptChan, 10)
	return TypeDefine.UserOpt{ID: u.ID, Ch: ch}
}
func GetViewDataCh(u FunctionalMember.FuncMember) TypeDefine.UserViewData {
	ch := make(TypeDefine.UserViewChan, 10)
	return TypeDefine.UserViewData{ID: u.ID, Ch: ch}
}

// UserRootHandle 用户业务
func UserRootHandle(user User.UserInf) error {
	//捕获异常
	defer Notice.RecoverPanic()

	if u, ok := user.(*FunctionalMember.FuncMember); !ok {
		return Error.ErrorInit("This is a guest", 400)
	} else {
		//
		params := GetUserParams(u)
		//启动业务逻辑
		ctx, cancel := context.WithTimeout(context.Background(), Define.TokenExpireDuration)
		Define.CancelMap[GetAccountNum(u.GetID())] = cancel
		go func() {
			defer cancel()
			err := Default.LoadUserRootRuleChain().Apply(ctx, params) //异步执行业务，业务结束则会cancel
			if err != nil {
				//用户出问题终止这个用户协程
				Error.NewErrHandle(err).WriteErr()
			} else {
				fmt.Println("用户业务结束")
			}
			if user.GetOnlineStatus() == true {
				user.ChangeOnlineStatus()
			}
		}()
		return nil
	}
}

// GetMsgReceiver 按用户消息的最后时间查找用户
func GetMsgReceiver(AccountNum int64) *FunctionalMember.FuncMember {
	return nil
}

// UserOnline 用户上线//返回用户id
func UserOnline(user User.UserInf) error {
	//捕获异常
	defer Notice.RecoverPanic()

	//改变在线状态
	if user.GetOnlineStatus() == false {
		return Error.ErrorInit("User online status is wrong", 400)
	}

	//start handler
	go func() {
		//捕获异常
		defer Notice.RecoverPanic()

		Error.NewErrHandle(UserRootHandle(user)).WriteErr().ViewErr()

	}()

	//日志
	NewLog := Log.NewUserLog(user.GetID(), time.Now(), User.ZeroTime.Add(time.Nanosecond))
	Error.NewErrHandle(Log.CreateUserLog(NewLog)).WriteErr()

	return nil
}

// UserOffline 用户下线
func UserOffline(user User.UserInf) error {
	//捕获异常
	defer Notice.RecoverPanic()

	//改变在线状态
	status := user.GetOnlineStatus()
	if status == false {
		return Error.ErrorInit("User online status is wrong", 400)
	}
	user.ChangeOnlineStatus()

	//日志
	ID := user.GetID()
	NewLog := Log.NewUserLog(ID, User.ZeroTime.Add(time.Nanosecond), time.Now())
	Error.NewErrHandle(Log.CreateUserLog(NewLog)).WriteErr()

	return nil
}
