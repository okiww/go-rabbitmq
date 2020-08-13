package mq

import (
	"github.com/spf13/cobra"
	"go-rabbitmq/cmd/mq/consumer"
	configs "go-rabbitmq/config"
	"go-rabbitmq/infra/mq"
)

var (
	consumerCMD  = &cobra.Command{
		Use:   "consumer",
		Short: "Example Consumer Message",
		Long:  "Example Consumer a message to mq",
		RunE:  runConsumer,
	}
)

func runConsumer(cmd *cobra.Command, args []string) error {
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

	// set config QOS
	mq.SetQOSCount(1)
	mq.SetQOSSize(1)
	mq.SetQOSGlobal(false)
	mq.QOS()

	p := consumer.NewConsumerMQ(mq)
	p.ConsumerMessage(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	mq.Close()
	return nil
}

// ServeMQConsumer return instance of serve mq command object
func ServeMQConsumer() *cobra.Command {
	return consumerCMD
}

