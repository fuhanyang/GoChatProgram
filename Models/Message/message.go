package Models

import (
	"MyTest/DAO/Mysql"
	"github.com/jinzhu/gorm"
)

type Message struct {
	tick       uint
	SenderID   uint
	ReceiverID uint
	Content    string
	gorm.Model
}

func HashTick(ID1 uint, ID2 uint) uint {
	t := ID1*100 + ID2
	return t
}

func NewMessage(senderID uint, receiverID uint, content string) *Message {
	t := HashTick(senderID, receiverID)

	return &Message{tick: t, Content: content, SenderID: senderID, ReceiverID: receiverID}
}

func WriteMsg(message *Message) error {
	Mysql.MysqlDb.Create(message)

	return nil
}

// 删
func DeleteMsg(msg Message) error {
	Mysql.MysqlDb.Delete(&msg)

	return nil
}

// 查
func GetMsgs(SeverID uint, ReciverID uint, limit uint) (M []Message, err error) {
	Mysql.MysqlDb.Order("created_at desc").Limit(limit).Find(&M, &Message{tick: HashTick(SeverID, ReciverID)})
	//
	return M, nil
}
