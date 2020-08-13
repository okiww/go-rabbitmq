package mq

import (
	"github.com/spf13/cobra"
	"go-rabbitmq/cmd/mq/publisher"

	//"go-rabbitmq/cmd/mq/publisher"
	configs "go-rabbitmq/config"
	mq "go-rabbitmq/infra/mq"
)

var (
	publisherCMD  = &cobra.Command{
		Use:   "publish",
		Short: "Example Publish Message",
		Long:  "Example publish a message to mq",
		RunE:  runPublisher,
	}
)

func runPublisher(cmd *cobra.Command, args []string) error {
	// initial config
	config := configs.InitConfig()

	// initial MessageQueue
	mq := mq.NewMQService(config.MQServer)
	// connect to rabbitMQ
	mq.Connect()
	// declare a queue
	q, _ := mq.QueueDeclare(
		"task_queue", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	// example publish message
	body := "Hello World!"
	p := publisher.NewPublisherMQ(mq)
	p.PublishMessage(
		"",
		q.Name,
		false,
		false,
		body,
	)

	mq.Close()
	return nil
}

// ServeMQPublisher return instance of serve mq command object
func ServeMQPublisher() *cobra.Command {
	return publisherCMD
}
