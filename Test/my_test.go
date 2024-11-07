package Test

import (
	"fmt"
	"reflect"
	"testing"
)

//
//func Test1(t *testing.T) {
//	defer rabbitmq.MqClose()
//	defer Mysql.MysqlClose()
//	rabbitmq.Init()
//	Mysql.Init()
//	Table.InitTable()
//	myopt := make(chan string, 10)
//	q, err := rabbitmq.Ch.QueueDeclare(
//		"sever",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("queue declare")
//	//rabbitmq.Ch.QueueBind(
//	//	q.Name, // queue name
//	//	"",     // routing key
//	//	"",     // exchange
//	//	false,
//	//	nil,
//	//)
//	msgs, _ := rabbitmq.Ch.Consume(
//		q.Name,
//		"",
//		true,
//		false,
//		false,
//		false,
//		nil)
//	go func() {
//		for d := range msgs {
//			fmt.Println(string(d.Body))
//			myopt <- string(d.Body)
//		}
//	}()
//	go func() {
//		router := gin.Default()
//		UsersGroup := router.Group("/users")
//		{
//			UsersGroup.GET(":accountNum", Routers.LoadUsersPage())
//		}
//		//查看别的用户
//		ReceiverGroup := UsersGroup.Group("/rec")
//		{
//			ReceiverGroup.GET("/:rec_id/:accountNum", Routers.LoadUserMsgPage())
//
//		}
//		// 使用 CORS 中间件
//		router.Use(func(c *gin.Context) {
//			c.Header("Access-Control-Allow-Origin", "*")                   // 允许所有域访问
//			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // 允许的 HTTP 方法
//			c.Header("Access-Control-Allow-Headers", "Content-Type")       // 允许的请求头
//
//			// 如果是 OPTIONS 请求，直接返回
//			if c.Request.Method == http.MethodOptions {
//				c.Status(http.StatusOK)
//				return
//			}
//
//			c.Next() // 继续处理请求
//		})
//
//		// 定义 /data 路由
//		router.GET("/data", handler())
//		router.Run(":9090")
//	}()
//	for {
//		s := <-myopt
//		switch s {
//		case "signup":
//			fmt.Println("user try to sign")
//			UserCreate.UserSignUp("123", "123")
//		case "signin":
//			account, _ := strconv.Atoi("117065")
//			pass := "123"
//			U, _ := UserCreate.UserSignIn(pass, account, "123")
//
//			UserHandle.UserOnline(U, context.Background())
//			TypeDefine.Mu.Lock()
//			obj := TypeDefine.UserMap[account]
//			TypeDefine.Mu.Unlock()
//
//			optmap := make(map[interface{}]interface{})
//			//optmap["opt_type"] = "check"
//			//optmap["id"] = uint(1)
//			obj.Opt.Ch <- optmap
//			time.Sleep(1 * time.Second)
//			optmap["opt_type"] = "sendMsg"
//			optmap["id"] = uint(1)
//			optmap["msg"] = <-myopt
//			obj.Opt.Ch <- optmap
//			time.Sleep(1 * time.Second)
//			optmap["id"] = <-myopt
//			optmap["opt_type"] = <-myopt
//			obj.Opt.Ch <- optmap
//
//		default:
//			fmt.Println("wrong")
//		}
//	}
//
//}

type a struct {
	name string
}
type b struct {
	a
}

func (A *a) checkT() {
	fmt.Println(reflect.TypeOf(A))
	A.name = "test"
}
func (A *a) GetName() {
	fmt.Println(A.name)
}

func (B *b) GetName() {
	fmt.Println(B.name)
}

func Test11(t *testing.T) {
	B := &b{a{"bName"}}
	B.GetName()
	B.checkT()
	B.GetName()
}
