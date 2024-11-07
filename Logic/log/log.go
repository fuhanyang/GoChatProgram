package log

import (
	"MyTest/Models/Log"
	"fmt"
	"log"
)

// SendErrToUser 向调用用户报告错误
func SendErrToUser(err string) {

}

func RecoverPanic() {
	if r := recover(); r != nil {
		err, ok := r.(string)
		if ok {
			fmt.Println(err)
			Log.CreateLog(Log.NewLog("Recovered from panic: " + err))
		} else {
			fmt.Println("Panic!!!")
			s := fmt.Sprintf("Recovered from panic: %v", r)
			fmt.Println(s)
			Log.CreateLog(Log.NewLog(s))
		}
		// 注意：如果这个函数是被defer调用的，并且你希望阻止panic继续传播，
		// 那么你只需要在这里处理panic（比如记录日志），而不需要返回任何值。
		// 如果需要向调用者报告错误，你可能需要设计一种不同的错误处理机制，
		// 因为panic一旦被recover，就不会再向上传播了。
		SendErrToUser(err)

	}
}

// WriteLog 日志写入
func WriteLog(TYPE string, msg string) error {
	MyLog := Log.NewLog(msg)
	log.Printf("TYPE :%s  message :%s ", TYPE, MyLog.Message)
	//

	return nil
}
