package Mq

import (
	"MyTest/Logic/Notice"
	"MyTest/Models/Error"
	"MyTest/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendNotice(notice string) error {
	//捕获异常
	defer Notice.RecoverPanic()

	if rabbitmq.Conn == nil {
		return Error.ErrorInit("not connected to RabbitMQ", 404)
	}
	//把公告发给用户
	err := rabbitmq.Ch.Publish(
		"notice_direct",
		"mq",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(notice),
		})
	if err != nil {
		return err
	}
	err = Notice.WriteLog("mq", notice)
	Error.FailOnError(err, "mq fail to write")
	return err
}
