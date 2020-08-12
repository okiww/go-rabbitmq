package mq

import (
	"github.com/streadway/amqp"
	"go-rabbitmq/config"
	"log"
)

type MQInterface interface {
	Connect() error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args map[string]interface{}) (amqp.Queue, error)
	Publish(exchanged, name string, mandatory, immediate bool, msg []byte) error
	Close() error
}

type mqService struct {
	MQConfig config.MQConfiguration
	conn     *amqp.Connection
	channel  *amqp.Channel
}

func (m *mqService) Close() error {
	if m.conn == nil {
		return nil
	}

	return m.conn.Close()
}

func (m *mqService) QueueDeclare(
	name string,
	durable,
	autoDelete,
	exclusive,
	noWait bool,
	args map[string]interface{},
) (amqp.Queue, error) {
	q, err := m.channel.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // no-wait
		args,       // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "failed to declare a queue", err)
		return q, err
	}
	return q, nil
}

func (m *mqService) Publish(exchanged, name string, mandatory, immediate bool, msg []byte) error {
	err := m.channel.Publish(
		exchanged, // exchange
		name,      // routing key
		mandatory, // mandatory
		immediate, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	log.Printf(" [x] Sent %s", msg)
	if err != nil {
		log.Fatalf("%s: %s", "failed to publish a message", err)
		return err
	}

	return nil
}

func (m *mqService) Connect() error {
	conn, err := amqp.Dial(m.MQConfig.Dial)
	if err != nil {
		log.Fatalf("%s: %s", "failed to connect rabbitMQ", err)
		return err
	}
	m.conn = conn

	ch, err := m.conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "failed to open a channel", err)
		return err
	}
	m.channel = ch

	log.Printf("success connect to rabbitMQ")
	return nil
}

func NewMQService(mqConfig config.MQConfiguration) MQInterface {
	return &mqService{
		MQConfig: mqConfig,
	}
}
