package User

import (
	"MyTest/DAO/Mysql"
	"MyTest/Models/Log"
	Models "MyTest/Models/Message"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	IP       string
	IsOnline bool
	IsGuest  bool
	gorm.Model
}
type UserInfo struct {
	Name     string `json:"username"`
	PassWord string `json:"password"`
}

type UserInf interface {
	UserSignUp() error
	UserSignIn() error
	UserSignOut() error
	GetIP() string
	GetOnlineStatus() bool
	GetID() uint
	GetUserType() bool
	ChangeName(string) error
	ChangeOnlineStatus() bool
}

var Guest = true
var FuncMember = false
var ZeroTime = time.Time{}

// 用户数据的各种行为
func (U *User) UserSignUp() error {
	//在数据库中创建用户信息
	Mysql.MysqlDb.Create(U)

	return nil
}

func (U *User) UserSignIn() error {
	U.IsOnline = true
	Mysql.MysqlDb.Model(U).Update("is_online", true)

	//获取上线时间写入数据库
	now := time.Now()
	newLog := Log.UserLog{ID: U.ID, Online: now, Offline: ZeroTime.Add(time.Nanosecond)}
	Mysql.MysqlDb.Save(&newLog)

	//如果成功上线了就返回true

	return nil
}

// 下线
func (U *User) UserSignOut() error {
	U.IsOnline = false
	Mysql.MysqlDb.Model(U).Update("is_online", false)

	//获取下线时间写入
	now := time.Now()
	log := Log.UserLog{ID: U.ID, Online: ZeroTime.Add(time.Nanosecond), Offline: now}
	Mysql.MysqlDb.Save(&log)

	//success
	return nil
}

func (U *User) GetIP() string {
	//

	return U.IP
}

func (U *User) GetOnlineStatus() bool {
	//

	return U.IsOnline
}

func (U *User) GetID() uint {
	return U.ID
}

func (U *User) GetUserType() bool {
	return U.IsGuest
}

func (U *User) SendMsgTo(rec UserInf, content string) error {
	msg := Models.NewMessage(U.ID, rec.GetID(), content)
	//
	Mysql.MysqlDb.Create(&msg)
	//

	return nil
}

func (U *User) ChangeOnlineStatus() bool {
	U.IsOnline = !U.IsOnline
	Mysql.MysqlDb.Model(U).Update("is_online", U.IsOnline)

	return true
}

func (U *User) ChangeIp(ip string) {
	U.IP = ip
	Mysql.MysqlDb.Model(U).Update("ip", U.IP)
}
