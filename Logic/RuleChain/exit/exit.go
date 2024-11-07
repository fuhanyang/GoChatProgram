package exit

import (
	"MyTest/Logic/RuleChain"
	"MyTest/Models/Error"
	"context"
)

var Exit = &Error.MyError{
	"Exit",
	1000,
}

type ExitRuleChain struct {
	RuleChain.BaseRuleChain
}

func NewExitRuleChain() RuleChain.RuleChain {
	return &ExitRuleChain{}
}

func (e *ExitRuleChain) Apply(ctx context.Context, p RuleChain.Params) error {
	return Exit
}
