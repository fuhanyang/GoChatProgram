package Load

import (
	"MyTest/Logic/Notice"
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	"context"
	"fmt"
)

type LoadFunctionRuleChain struct {
	RuleChain.BaseRuleChain
}

func NewLoadFunctionRuleChain(next RuleChain.RuleMap) RuleChain.RuleChain {
	return &LoadFunctionRuleChain{
		BaseRuleChain: RuleChain.BaseRuleChain{
			NextRule: next,
		},
	}
}
func (L *LoadFunctionRuleChain) Apply(ctx context.Context, params RuleChain.Params) error {

	//捕获异常
	defer Notice.RecoverPanic()

	for {
		//view.PrintUser()
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if ch, ok := params[TypeDefine.Ch].(TypeDefine.UserOptChan); !ok {
			return Error.ErrorInit("Get User option chan wrong!", 400)
		} else {
			//启动业务
			for {
				fmt.Println("load chain")
				var opt map[interface{}]interface{}
				select {
				case <-ctx.Done():
					return ctx.Err()
				case opt = <-ch:
				}

				params[TypeDefine.Opt] = opt
				fmt.Println(opt)
				err := L.ApplyNext(ctx, params) //调用后续业务
				//有错误则打印错误，继续循环,exit退出
				if err != nil {
					Error.NewErrHandle(err).WriteErr()
					return nil
				}

			}

		}

	}
}
