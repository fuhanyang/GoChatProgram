package RuleChain

import (
	"MyTest/Logic/user_managment/TypeDefine"
	"MyTest/Models/Error"
	"context"
)

const (
	TurnOff   = "TurnOff"
	Opt_Type  = "Opt_Type"
	SendMsg   = "SendMsg"
	CheckUser = "CheckUser"
	SeekUser  = "SeekUser"
	EXIT      = "Exit"
)

type Params map[interface{}]interface{} //用户自己的参数配置
type RuleMap map[interface{}]RuleChain

type RuleChain interface {
	Apply(ctx context.Context, p Params) error
	Next(param interface{}) RuleChain
}

type BaseRuleChain struct {
	NextRule RuleMap
}

func (b *BaseRuleChain) Apply(ctx context.Context, p Params) error {
	panic("not implement")
}

func (b *BaseRuleChain) Next(param interface{}) RuleChain {
	return b.NextRule[param]
}
func (b *BaseRuleChain) ApplyNext(ctx context.Context, params Params) error {
	//用户操作来调用相应业务
	opt, ok := params[TypeDefine.Opt].(map[interface{}]interface{})
	if !ok {
		return Error.ErrorInit("opt get wrong", 400)
	}
	if opt[Opt_Type] != nil {
		if b.Next(opt[Opt_Type]) != nil {
			return b.Next(opt[Opt_Type]).Apply(ctx, params)
		} else {
			//调用不存在的业务
			return Error.ErrorInit("Handle does not exist", 400)
		}
	} else {
		return Error.ErrorInit("Null type option", 400)
	}
}
