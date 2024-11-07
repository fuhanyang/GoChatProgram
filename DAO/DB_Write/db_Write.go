package DB_Write

import (
	"MyTest/Models/Error"
)

func LogErr(err error) {
	if err != nil {
		Error.NewErrHandle(err).WriteErr()
	}
}
