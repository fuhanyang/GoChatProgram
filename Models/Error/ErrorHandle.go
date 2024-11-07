package Error

import (
	"MyTest/Models/Log"
	"fmt"
)

type ErrHandle struct {
	err error
}

func NewErrHandle(err error) *ErrHandle {
	return &ErrHandle{err: err}
}

func (e *ErrHandle) WriteErr() *ErrHandle {
	if e != nil {
		err := e.err
		if err != nil {
			fmt.Println(err.Error())
			Log.CreateLog(Log.NewLog("WriteErr: " + err.Error()))
			return e
		}
	}
	return nil
}

func (e *ErrHandle) ViewErr() *ErrHandle {

	if e != nil {
		err := e.err
		if err != nil {

		}
	}
	return e
}
