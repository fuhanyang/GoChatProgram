package UserRuleChain

import "C"
import (
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/message_systerm"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	"context"
	"errors"
	"fmt"
)

type DialogRuleChain struct {
	RuleChain.BaseRuleChain
}

func NewDialogRuleChain(next RuleChain.RuleMap) RuleChain.RuleChain {
	return &DialogRuleChain{
		BaseRuleChain: RuleChain.BaseRuleChain{
			NextRule: next,
		},
	}
}
func (D *DialogRuleChain) Apply(ctx context.Context, params RuleChain.Params) error {
	if ch, ok := params["ch"].(TypeDefine.UserOptChan); !ok {
		return Error.ErrorInit("Get User option chan wrong!", 400)
	} else {
		opt := params["opt"].(map[interface{}]interface{})
		id, ok1 := opt["id"].(uint)
		my_id, ok2 := params["id"].(uint)
		msg, ok3 := opt["msg"].(string)
		if !ok1 || !ok2 || !ok3 {
			return Error.ErrorInit("Wrong data type [dialog]", 400)
		}
		Error.NewErrHandle(message_systerm.SendMsg(my_id, id, msg)).WriteErr().ViewErr()
		//后续业务
		for {
			opt = <-ch
			params["opt"] = opt
			fmt.Println(opt)
			err := D.ApplyNext(ctx, params)
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
