package consumer

import (
	"go-rabbitmq/infra/mq"
	"log"
)

type Consumer interface {
	ConsumerMessage(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{})
}

type consumer struct {
	mq.MQInterface
}

func (c consumer) ConsumerMessage(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{}) {
	forever := c.Consumer(
		queue,
		consumer,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func NewConsumerMQ(mqInterface mq.MQInterface) Consumer {
	return consumer{
		mqInterface,
	}
}
