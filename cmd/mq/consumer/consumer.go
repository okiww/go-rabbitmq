package consumer

import "go-rabbitmq/infra/mq"

type Consumer interface {
	ConsumerMessage(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{})
}

type consumer struct {
	mq.MQInterface
}

func (c consumer) ConsumerMessage(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{}) {
	c.Consumer(
		queue,
		consumer,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)
}

func NewConsumerMQ(mqInterface mq.MQInterface) Consumer {
	return consumer{
		mqInterface,
	}
}
