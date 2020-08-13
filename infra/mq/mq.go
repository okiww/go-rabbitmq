package mq

import (
	"bytes"
	"github.com/streadway/amqp"
	"go-rabbitmq/config"
	"log"
	"time"
)

type Type struct {}

type MQInterface interface {
	// MQ Method
	Connect() error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args map[string]interface{}) (amqp.Queue, error)
	Publisher(exchanged, name string, mandatory, immediate bool, msg []byte) error
	Consumer(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{}) chan bool
	Close() error

	// QOS Method
	SetQOSCount(i int) ConfigFunc
	SetQOSSize(size int) ConfigFunc
	SetQOSGlobal(global bool) ConfigFunc
	QOS() error
}

type mqService struct {
	MQConfig config.MQConfiguration
	conn     *amqp.Connection
	channel  *amqp.Channel
	QOSConf  ConfigQOS
}

func (m *mqService) QOS() error {
	err := m.channel.Qos(
		m.QOSConf.Count,
		m.QOSConf.Size,
		m.QOSConf.Global,
	)

	if err != nil {
		log.Fatalf("%s: %s", "failed to set QOS", err)
		return err
	}
	log.Printf("success set QOS")
	return nil
}

func (m *mqService) Consumer(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{}) chan bool{
	msgs, err := m.channel.Consume(
		queue,
		consumer,
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)
	if err != nil {
		log.Fatalf("%s: %s", "failed to register consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	return forever
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

func (m *mqService) Publisher(exchanged, name string, mandatory, immediate bool, msg []byte) error {
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
