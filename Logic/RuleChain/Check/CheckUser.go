package Check

import (
	"MyTest/Logic/Notice"
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	Models "MyTest/Models/Message"
	"MyTest/view"
	"context"
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
	defer Notice.RecoverPanic()

	if ch, ok := params[TypeDefine.Ch].(TypeDefine.UserOptChan); !ok {
		return Error.ErrorInit("Get User option chan wrong!", 400)
	} else {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		//当前的业务调用
		opt, _ := params[TypeDefine.Opt].(map[interface{}]interface{})
		id, ok1 := opt[TypeDefine.ID].(uint)
		my_id, ok2 := params[TypeDefine.ID].(uint)
		if !ok1 || !ok2 {
			return Error.ErrorInit("Wrong data type [check]", 400)
		}
		for {
			view.PrintMsg(CheckUserMsg(id, my_id), params[TypeDefine.AccountNum].(int64))
			//递归调用之后的业务
			select {
			case <-ctx.Done():
				return ctx.Err()
			case opt = <-ch:
			}
			fmt.Println("this opt:", "check", "next opt:", opt)
			params[TypeDefine.Opt] = opt
			err := C.ApplyNext(ctx, params)

			//有错误则打印错误，继续循环,exit退出
			if err != nil {
				Error.NewErrHandle(err).WriteErr().ViewErr()
				return nil
			}
		}
	}
}
func CheckUserMsg(ID uint, MyID uint) []Models.Message {
	//捕获异常
	defer Notice.RecoverPanic()

	msg, err := Models.GetMsgs(MyID, ID, 10)
	Error.NewErrHandle(err).WriteErr().ViewErr()

	return msg

}
func GetUserCheckOpt(accountNum int64, rec_id string) {
	UserObj := TypeDefine.UserMap[accountNum]

	recNum, err := strconv.Atoi(rec_id)
	Error.NewErrHandle(err).WriteErr().ViewErr()
	UserObj.Opt.Ch <- map[interface{}]interface{}{
		RuleChain.Opt_Type: RuleChain.CheckUser,
		TypeDefine.ID:      uint(recNum),
	}
	//
}
