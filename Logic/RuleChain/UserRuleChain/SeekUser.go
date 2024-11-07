package UserRuleChain

import "C"
import (
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/log"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	"MyTest/Models/Users/FunctionalMember"
	"MyTest/view"
	"context"
	"errors"
)

type SeekRuleChain struct {
	RuleChain.BaseRuleChain
}

func NewSeekRuleChain(next RuleChain.RuleMap) RuleChain.RuleChain {
	return &SeekRuleChain{
		BaseRuleChain: RuleChain.BaseRuleChain{
			NextRule: next,
		},
	}
}

func (S *SeekRuleChain) Apply(ctx context.Context, params RuleChain.Params) error {
	if ch, ok := params["ch"].(TypeDefine.UserOptChan); !ok {
		return Error.ErrorInit("Get User option chan wrong!", 400)
	} else {
		//当前的业务调用
		opt, _ := params["opt"].(map[interface{}]interface{})
		name, _ := opt["name"].(string)
		M := SeekUser(name, 10)

		for {
			view.PrintUser(M)
			//递归调用之后的业务
			opt = <-ch
			params["opt"] = opt
			err := S.ApplyNext(ctx, params)
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

func SeekUser(name string, limit int) []FunctionalMember.FuncMember {
	//捕获异常
	defer log.RecoverPanic()

	var M []FunctionalMember.FuncMember
	M, err := FunctionalMember.GetUsers(limit, name)
	Error.NewErrHandle(err).WriteErr().ViewErr()

	return M
}
