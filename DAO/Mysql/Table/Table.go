package Table

import (
	"MyTest/DAO/Mysql"
	"MyTest/Models/Log"
	Models "MyTest/Models/Message"
	"MyTest/Models/Users/FunctionalMember"
	"MyTest/Models/Users/Guest"
	"fmt"
)

func InitTable() {

	Mysql.MysqlDb.AutoMigrate(&Models.Message{})
	Mysql.MysqlDb.AutoMigrate(&FunctionalMember.FuncMember{})
	Mysql.MysqlDb.AutoMigrate(&Guest.Guest{})
	Mysql.MysqlDb.AutoMigrate(&Log.UserLog{})
	Mysql.MysqlDb.AutoMigrate(&Log.Log{})
	//
	fmt.Println("table auto migrate")
}
