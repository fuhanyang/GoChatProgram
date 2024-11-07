package UserRuleChain

import "C"
import (
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/log"
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	"context"
	"errors"
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
	defer log.RecoverPanic()

	for {
		//view.PrintUser()

		if ch, ok := params["ch"].(TypeDefine.UserOptChan); !ok {
			return Error.ErrorInit("Get User option chan wrong!", 400)
		} else {
			//启动业务
			for {
				fmt.Println("load chain")
				opt := <-ch
				params["opt"] = opt
				fmt.Println(opt)
				if opt_type, ok := opt["opt_type"].(string); !ok {
					Error.NewErrHandle(Error.ErrorInit("Get User option type wrong!", 400)).WriteErr().ViewErr()
				} else {
					params[opt_type] = opt_type
					err := L.ApplyNext(ctx, params)
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

	}
}
