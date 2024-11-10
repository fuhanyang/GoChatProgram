package Dialog

import (
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/message_systerm"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	"context"
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
	if ch, ok := params[TypeDefine.Ch].(TypeDefine.UserOptChan); !ok {
		return Error.ErrorInit("Get User option chan wrong!", 400)
	} else {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		opt := params[TypeDefine.Opt].(map[interface{}]interface{})
		id, ok1 := opt["rec_id"].(uint)
		my_id, ok2 := params[TypeDefine.ID].(uint)
		msg, ok3 := opt[message_systerm.Msg].(string)

		if !ok1 || !ok2 || !ok3 {
			return Error.ErrorInit("Wrong data type [dialog]", 400)
		}

		//
		Error.NewErrHandle(message_systerm.SendMsg(my_id, id, msg)).WriteErr().ViewErr()

		//后续业务
		for {
			var opt map[interface{}]interface{}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case opt = <-ch:
			}
			params[TypeDefine.Opt] = opt
			fmt.Println(opt)
			err := D.ApplyNext(ctx, params)
			//有错误则打印错误，继续循环,exit退出
			if err != nil {
				Error.NewErrHandle(err).WriteErr().ViewErr()
				return nil
			}
		}
	}

}
