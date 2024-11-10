package Seek

import (
	"MyTest/Logic/Notice"
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	"MyTest/Models/Users/FunctionalMember"
	"MyTest/view"
	"context"
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
	if ch, ok := params[TypeDefine.Ch].(TypeDefine.UserOptChan); !ok {
		return Error.ErrorInit("Get User option chan wrong!", 400)
	} else {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		//当前的业务调用
		opt, _ := params[TypeDefine.Opt].(map[interface{}]interface{})
		name, _ := opt[TypeDefine.Name].(string)
		M := SeekUser(name, 10)

		for {
			view.PrintUser(M)
			//递归调用之后的业务
			var opt map[interface{}]interface{}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case opt = <-ch:
			}
			params[TypeDefine.Opt] = opt
			err := S.ApplyNext(ctx, params)
			//有错误则打印错误，继续循环,exit退出
			if err != nil {
				Error.NewErrHandle(err).WriteErr().ViewErr()
				return nil
			}
		}
	}

}

func SeekUser(name string, limit int) []FunctionalMember.FuncMember {
	//捕获异常
	defer Notice.RecoverPanic()

	var M []FunctionalMember.FuncMember
	M, err := FunctionalMember.GetUsers(limit, name)
	Error.NewErrHandle(err).WriteErr().ViewErr()

	return M
}
