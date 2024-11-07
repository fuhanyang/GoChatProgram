package Log

import (
	"MyTest/DAO/Mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type UserLog struct {
	Index   uint `gorm:"primary_key"`
	ID      uint
	Online  time.Time
	Offline time.Time
}

// 增
func NewUserLog(ID uint, OnlineTime time.Time, OfflineTime time.Time) *UserLog {
	return &UserLog{ID: ID, Online: OnlineTime, Offline: OfflineTime}
}

func CreateUserLog(log *UserLog) error {
	Mysql.MysqlDb.Create(log)

	return nil
}

// 查
func GetUserLog(ID uint) (L *UserLog, err error) {
	Mysql.MysqlDb.First(L, ID)
	//

	return L, nil
}

// 删
func DeleteUserLog(ID uint) error {
	Mysql.MysqlDb.Delete(&UserLog{ID: ID})

	//

	return nil
}

type Log struct {
	Message string
	gorm.Model
}

// 增
func NewLog(message string) *Log {
	return &Log{Message: message}
}

func CreateLog(log *Log) error {

	Mysql.MysqlDb.Create(log)

	return nil
}

// 查
func GetLog(ID uint) (L *Log, err error) {
	Mysql.MysqlDb.First(L, ID)
	//

	return L, nil
}

// 删
func DeleteLog(ID uint) error {

	Mysql.MysqlDb.Delete(&Log{
		Model: gorm.Model{ID: ID},
	})

	//

	return nil
}
