package Default

import (
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/RuleChain/Check"
	"MyTest/Logic/RuleChain/Dialog"
	"MyTest/Logic/RuleChain/Load"
	"MyTest/Logic/RuleChain/Seek"
	"MyTest/Logic/RuleChain/exit"
	"fmt"
)

// LoadUserRootRuleChain 责任链的加载
func LoadUserRootRuleChain() (RootRule RuleChain.RuleChain) {
	exitRule := exit.NewExitRuleChain()

	//只实现发消息的逻辑就可以，用mq发，并获取响应，收的逻辑写在客户端，接mq相应就行
	dialogRule := Dialog.NewDialogRuleChain(func() (next RuleChain.RuleMap) {
		next = make(RuleChain.RuleMap)
		next[RuleChain.TurnOff] = exitRule
		return next
	}())
	checkRule := Check.NewCheckRuleChain(func() (next RuleChain.RuleMap) {
		next = make(RuleChain.RuleMap)
		next[RuleChain.SendMsg] = dialogRule
		next[RuleChain.TurnOff] = exitRule
		return next
	}())

	seekRule := Seek.NewSeekRuleChain(func() (next RuleChain.RuleMap) {
		next = make(RuleChain.RuleMap)
		next[RuleChain.CheckUser] = checkRule
		next[RuleChain.TurnOff] = exitRule
		return next
	}())

	loadFunctionRule := Load.NewLoadFunctionRuleChain(func() (next RuleChain.RuleMap) {
		next = make(RuleChain.RuleMap)
		next[RuleChain.SeekUser] = seekRule
		next[RuleChain.CheckUser] = checkRule
		next[RuleChain.TurnOff] = exitRule
		return next
	}())

	if loadFunctionRule == nil {
		fmt.Println("LoadFunctionRule is nil!!!!")
	}
	return loadFunctionRule
}
