package Mysql

import (
	"MyTest/Settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MysqlDb *gorm.DB

func Init(config *Settings.MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", config.UserName, config.Password, config.Host, config.Port, config.DbName)
	var err error
	MysqlDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	MysqlDb.DB().SetMaxIdleConns(10)
	MysqlDb.DB().SetMaxOpenConns(100)
	fmt.Println("mysql connect success")
	return MysqlDb, nil

	//var err error
	//dsn := "root:56563096660fc@tcp(127.0.0.1:3306)/sever_demo?charset=utf8&parseTime=True&loc=Local"
	//MysqlDb, err = gorm.Open("mysql", dsn)
	//return MysqlDb, err
}
func MysqlClose() error {
	return MysqlDb.Close()

}
