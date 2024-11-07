package Error

import (
	"fmt"
	"log"
)

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code:%d, message:%s", e.Code, e.Message)
}

// ErrorInit 记录某个错误
func ErrorInit(msg string, code int) error {
	return &MyError{msg, code}
}

// FailOnError 记录错误日志并panic错误
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
