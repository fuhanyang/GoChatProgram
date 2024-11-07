package FunctionalMember

import (
	"MyTest/DAO/Mysql"
	"MyTest/Models/Error"
	"MyTest/Models/Log"
	Models "MyTest/Models/Message"
	"MyTest/Models/Users/User"
	"fmt"
	"time"
)

type FuncMember struct {
	User.User
	AccountNum int64 `gorm:"primary key" json:"account_num"`
	User.UserInfo
}

func NewFuncMember(Name string, IP string, AccountNum int64, PassWord string) *FuncMember {
	M := &FuncMember{
		User: User.User{
			IP:       IP,
			IsOnline: false,
			IsGuest:  User.FuncMember,
		},
		AccountNum: AccountNum,
		UserInfo: User.UserInfo{
			Name:     Name,
			PassWord: PassWord,
		},
	}

	fmt.Println("new m")
	return M
}

// 用户数据的各种行为

func (M *FuncMember) UserSignUp() error {
	//在数据库中创建用户信息
	Mysql.MysqlDb.Create(M)
	fmt.Println("write M:", M)
	return nil
}

func (M *FuncMember) UserSignIn() error {
	M.IsOnline = true
	Mysql.MysqlDb.Model(M).Update("is_online", true)

	//获取上线时间写入数据库
	now := time.Now()
	newLog := Log.UserLog{ID: M.ID, Online: now, Offline: User.ZeroTime.Add(time.Nanosecond)}
	Mysql.MysqlDb.Save(&newLog)

	//如果成功上线了就返回true

	return nil
}

// UserSignOut 用户下线
func (M *FuncMember) UserSignOut() error {
	M.IsOnline = false
	Mysql.MysqlDb.Model(M).Update("is_online", false)

	//获取下线时间写入
	now := time.Now()
	log := Log.UserLog{ID: M.ID, Online: User.ZeroTime.Add(time.Nanosecond), Offline: now}
	Mysql.MysqlDb.Save(&log)

	//success
	return nil
}

func (M *FuncMember) SendMsgTo(rec User.UserInf, content string) error {
	msg := Models.NewMessage(M.ID, rec.GetID(), content)
	//
	Mysql.MysqlDb.Create(&msg)
	//

	return nil
}

func (M *FuncMember) ChangeOnlineStatus() bool {
	M.IsOnline = !M.IsOnline
	Mysql.MysqlDb.Model(M).Update("is_online", M.IsOnline)

	return true
}

func (M *FuncMember) ChangeIp(ip string) {
	M.IP = ip
	Mysql.MysqlDb.Model(M).Update("ip", M.IP)
}

// UserLogOff 注销
func (M *FuncMember) UserLogOff() bool {
	M.IsOnline = false

	Mysql.MysqlDb.Delete(M)
	now := time.Now()
	log := Log.UserLog{ID: M.ID, Online: User.ZeroTime.Add(time.Nanosecond), Offline: now}
	Mysql.MysqlDb.Delete(&log)

	return true
}

func GetFuncMember(ID uint) (M *FuncMember, err error) {
	Mysql.MysqlDb.First(M, ID)
	//

	return M, nil
}
func GetUsers(limit int, name string) (M []FuncMember, err error) {
	Mysql.MysqlDb.Order("create_at desc").Limit(limit).Find(&M, &FuncMember{UserInfo: User.UserInfo{Name: name}})
	//
	return M, nil
}

// 改
func (M *FuncMember) ChangePassWord(s string) error {
	if M == nil {
		return Error.ErrorInit("Empty data of user", 400)

	}

	M.PassWord = s
	Mysql.MysqlDb.Model(M).Update("PassWord", s)

	return nil
}
func (M *FuncMember) ChangeName(s string) error {
	if M == nil {
		return Error.ErrorInit("Empty data of user", 400)
	}

	M.Name = s
	Mysql.MysqlDb.Model(M).Update("Name", s)

	return nil
}

// 查
func (M *FuncMember) GetPassWord() (string, error) {
	//

	return M.PassWord, nil
}
func (M *FuncMember) GetName() (string, error) {
	//

	return M.Name, nil
}

func GetTheData(ID uint) (M *FuncMember, err error) {
	Mysql.MysqlDb.First(M, ID)
	//

	return M, nil
}
