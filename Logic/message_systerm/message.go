package message_systerm

import (
	"MyTest/Models/Error"
	Models "MyTest/Models/Message"
	"MyTest/rabbitmq"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"strconv"
)

// SendMsg 用消息队列按时间发给用户
func SendMsg(SenderID uint, ReceiverID uint, msg string) error {
	fmt.Println(SenderID, "send msg to", ReceiverID)

	err := Models.WriteMsg(Models.NewMessage(SenderID, ReceiverID, msg))
	Error.NewErrHandle(err).WriteErr()

	id := strconv.Itoa(int(ReceiverID))

	err = rabbitmq.Ch.Publish(
		"user_msg_direct",
		id,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})

	return err
}

func ShowMsg() {

}
