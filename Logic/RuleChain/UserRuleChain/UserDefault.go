package UserRuleChain

import (
	"MyTest/Logic/RuleChain"
	"MyTest/Logic/RuleChain/exit"
	"fmt"
)

// LoadUserRootRuleChain 责任链的加载
func LoadUserRootRuleChain() (RootRule RuleChain.RuleChain) {
	exitRule := exit.NewExitRuleChain()

	//只实现发消息的逻辑就可以，用mq发，并获取响应，收的逻辑写在客户端，接mq相应就行
	dialogRule := NewDialogRuleChain(func() (next RuleChain.RuleMap) {
		next = make(RuleChain.RuleMap)
		next["exit"] = exitRule
		return next
	}())
	checkRule := NewCheckRuleChain(func() (next RuleChain.RuleMap) {
		next = make(RuleChain.RuleMap)
		next["sendMsg"] = dialogRule
		next["exit"] = exitRule
		return next
	}())

	seekRule := NewSeekRuleChain(func() (next RuleChain.RuleMap) {
		next = make(RuleChain.RuleMap)
		next["check"] = checkRule
		next["exit"] = exitRule
		return next
	}())

	loadFunctionRule := NewLoadFunctionRuleChain(func() (next RuleChain.RuleMap) {
		next = make(RuleChain.RuleMap)
		next["seek"] = seekRule
		next["check"] = checkRule
		next["exit"] = exitRule
		return next
	}())

	if loadFunctionRule == nil {
		fmt.Println("LoadFunctionRule is nil!!!!")
	}
	return loadFunctionRule
}
