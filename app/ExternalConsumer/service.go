package ExternalConsumer

import (
	"log"
	"time"

	"github.com/streadway/amqp"
	client "github.com/wade-sam/fyp-backup-server/pkg/Client"
	//"github.com/wade-sam/fyp-backup-server/pkg/client"
)

//**Route-key: Type.Policy.DeviceName

type ConsumerConfig struct {
	ExchangeName string
	ExchangeType string
	RoutingKey   string
	QueueName    string
	ConsumerName string
	Reconnect    struct {
		MaxAttempt int
		Interval   time.Duration
	}
}

func Initialise(clientService client.ExternalClientService, connection *amqp.Connection) {

	consumerconfig := ConsumerConfig{
		ExchangeName: "main",
		ExchangeType: "topic",
		RoutingKey:   "#.backupserver",
		QueueName:    "backupserver",
		ConsumerName: "backupserver",
	}
	consumerconfig.Reconnect.MaxAttempt = 60
	consumerconfig.Reconnect.Interval = 1 * time.Second

	consumerInstance := NewConsumer(consumerconfig, connection)
	if err := consumerInstance.Start(); err != nil {
		log.Fatalln("Unable to start consumer", err)
	}
	select {}
}

type Consumer struct {
	config     ConsumerConfig
	connection *amqp.Connection
}

func NewConsumer(config ConsumerConfig, conn *amqp.Connection) *Consumer {
	return &Consumer{
		config:     config,
		connection: conn,
	}
}

func (c *Consumer) Start() error {
	chn, err := c.connection.Channel()
	if err != nil {
		return err
	}

	if err := chn.ExchangeDeclare(
		c.config.ExchangeName,
		c.config.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil
	}
	if _, err := chn.QueueDeclare(
		c.config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil
	}

	if err := chn.QueueBind(
		c.config.QueueName,
		c.config.RoutingKey,
		c.config.ExchangeName,
		false,
		nil,
	); err != nil {
		return err
	}
	go c.consume(chn)
	return nil
}

func (c *Consumer) consume(channel *amqp.Channel) {
	msgs, err := channel.Consume(
		c.config.QueueName,
		c.config.ConsumerName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Could not start consumer", err)
		return
	}

	for msg := range msgs {
		log.Println("Consumed", string(msg.Body))
		if err := msg.Ack(false); err != nil {
			log.Println("unable to acknowledge the message, dropped", err)
		}
	}

	log.Println("Exiting")

}
