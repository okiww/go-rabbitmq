package publisher

import (
	"github.com/google/martian/log"
	"go-rabbitmq/infra/mq"
)

type Publish interface {
	PublishMessage(exchanged, name string, mandatory, immediate bool, msg string)
}

type publisher struct {
	mq.MQInterface
}

func (p publisher) PublishMessage(exchanged, name string, mandatory, immediate bool, msg string) {
	err := p.Publisher(
		exchanged,
		name,
		mandatory,
		immediate,
		[]byte(msg),
	)
	if err != nil {
		log.Errorf(err.Error())
	}
}

func NewPublisherMQ(mqInterface mq.MQInterface) Publish {
	return publisher{
		mqInterface,
	}
}
