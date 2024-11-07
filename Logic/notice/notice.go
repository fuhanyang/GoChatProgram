package notice

import (
	"MyTest/Logic/log"
	"MyTest/Models/Error"
	"MyTest/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendNotice(notice string) error {
	//捕获异常
	defer log.RecoverPanic()

	if rabbitmq.Conn == nil {
		return Error.ErrorInit("not connected to RabbitMQ", 404)
	}
	//把公告发给用户
	err := rabbitmq.Ch.Publish(
		"notice_direct",
		"notice",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(notice),
		})
	if err != nil {
		return err
	}
	err = log.WriteLog("notice", notice)
	Error.FailOnError(err, "notice fail to write")
	return err
}
