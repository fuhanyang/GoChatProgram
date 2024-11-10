package main

import (
	"MyTest/DAO/Mysql"
	"MyTest/DAO/Mysql/Table"
	"MyTest/Logic/user_managment/Snowflake"
	"MyTest/Routers"
	"MyTest/Settings"
	"fmt"
)

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
func main() {
	err := Settings.Init()
	if err != nil {
		panic(err)
	}

	//rabbitmq
	//rabbitmq.Init(Settings.Config.RabbitMQConfig)
	//defer rabbitmq.MqClose()

	//mysql
	Mysql.Init(Settings.Config.MysqlConfig)
	defer Mysql.MysqlClose()
	Table.InitTable()

	//redis

	//Redis.Init(Settings.Config.RedisConfig)

	//router
	router := Routers.Init()
	if router == nil {
		panic("router is nil")
	}

	//snowflake
	if err = Snowflake.Init(Settings.Config.SnowFlakeConfig); err != nil {
		fmt.Println(err)
	}

	router.Run(":9090")

}
