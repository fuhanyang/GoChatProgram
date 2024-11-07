package UserRuleChain

import (
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/log"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	Models "MyTest/Models/Message"
	"MyTest/view"
	"context"
	"errors"
	"fmt"
	"strconv"
)

type CheckRuleChain struct {
	RuleChain.BaseRuleChain
}

func NewCheckRuleChain(next RuleChain.RuleMap) RuleChain.RuleChain {
	return &CheckRuleChain{
		BaseRuleChain: RuleChain.BaseRuleChain{
			NextRule: next,
		},
	}
}
func (C *CheckRuleChain) Apply(ctx context.Context, params RuleChain.Params) error {
	fmt.Println("check handle")

	//捕获异常
	defer log.RecoverPanic()

	if ch, ok := params["ch"].(TypeDefine.UserOptChan); !ok {
		return Error.ErrorInit("Get User option chan wrong!", 400)
	} else {
		//当前的业务调用
		opt, _ := params["opt"].(map[interface{}]interface{})
		id, ok1 := opt["id"].(uint)
		my_id, ok2 := params["id"].(uint)
		if !ok1 || !ok2 {
			return Error.ErrorInit("Wrong data type [check]", 400)
		}
		for {
			view.PrintMsg(CheckUserMsg(id, my_id), params["accountNum"].(int64))
			//递归调用之后的业务
			opt = <-ch
			fmt.Println("this opt:", "check", "next opt:", opt)
			params["opt"] = opt
			err := C.ApplyNext(ctx, params)

			//有错误则打印错误，继续循环,exit退出
			if err == nil {
				continue
			}
			var err1 *Error.MyError
			if errors.As(err, &err1) && err1.Code == 1000 {
				//exit
				return nil
			}
			Error.NewErrHandle(err1).WriteErr().ViewErr()

		}
	}
}
func CheckUserMsg(ID uint, MyID uint) []Models.Message {
	//捕获异常
	defer log.RecoverPanic()

	msg, err := Models.GetMsgs(MyID, ID, 10)
	Error.NewErrHandle(err).WriteErr().ViewErr()

	return msg

}
func GetUserCheckOpt(accountNum int64, rec_id string) {
	UserObj := TypeDefine.UserMap[accountNum]

	recNum, err := strconv.Atoi(rec_id)
	Error.NewErrHandle(err).WriteErr().ViewErr()
	UserObj.Opt.Ch <- map[interface{}]interface{}{
		"opt_type": "check",
		"id":       uint(recNum),
	}
	//
}
