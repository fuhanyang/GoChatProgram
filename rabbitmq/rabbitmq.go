package rabbitmq

import (
	"MyTest/Models/Error"
	"MyTest/Settings"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"strconv"
)

var Conn *amqp.Connection
var Ch *amqp.Channel

func MqClose() error {
	err := Ch.Close()
	if err != nil {
		return err
	}
	err = Conn.Close()

	return err
}

func Init(config *Settings.RabbitMQConfig) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.UserName, config.Password, config.Host, config.Port)
	coon, err := amqp.Dial(dsn)
	Error.FailOnError(err, "Failed to connect to RabbitMQ")

	Ch, err = coon.Channel()
	Error.FailOnError(err, "Failed to open a Channel")

	err = Ch.ExchangeDeclare(
		"notice_direct",
		"direct",
		false,
		false,
		false,
		false,
		nil)
	Error.FailOnError(err, "Failed to declare an exChange")

	err = Ch.ExchangeDeclare(
		"user_msg_direct",
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	fmt.Println("Mq connect")
}

func UserMsgQueueDeclare(reciver_id uint, name string) {
	reciver_id = 1
	name = "123"
	q, err := Ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	Error.NewErrHandle(err).WriteErr()
	id := strconv.Itoa(int(reciver_id))
	err = Ch.QueueBind(
		q.Name,            // queue name
		id,                // routing key
		"user_msg_direct", // exchange
		false,
		nil,
	)
	Error.NewErrHandle(err).WriteErr()

}
