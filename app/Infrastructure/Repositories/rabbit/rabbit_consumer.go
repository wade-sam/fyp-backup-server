package rabbit

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
	"github.com/wade-sam/fyp-backup-server/entity"
)

type ConsumerConfig struct {
	ExchangeName string
	ExchangeType string
	RoutingKey   string
	QueueName    string
	ConsumerName string
	MaxAttempt   int
	Interval     time.Duration
	connection   *amqp.Connection
}

func (b *Broker) Start() (*amqp.Channel, error) {
	con, err := b.Connection()
	if err != nil {
		return nil, err
	}
	chn, err := con.Channel()
	if err != nil {
		return nil, err
	}

	if err := chn.ExchangeDeclare(
		b.Consumer.ExchangeName,
		b.Consumer.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}
	if _, err := chn.QueueDeclare(
		b.Consumer.QueueName,
		true,
		false,
		false,
		false,
		amqp.Table{"x-message-ttl": 6000},
	); err != nil {
		return nil, err
	}

	if err := chn.QueueBind(
		b.Consumer.QueueName,
		b.Consumer.RoutingKey,
		b.Consumer.ExchangeName,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	return chn, nil
}

func (b *Broker) Consume(channel *amqp.Channel) error {
	msgs, err := channel.Consume(
		b.Consumer.QueueName,
		b.Consumer.ConsumerName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		var d entity.Directory
		var file entity.ClientFile
		var holder entity.ClientFileHolder
		var s string
		switch msg.Type {

		case "backup":
			var f entity.File
			json.Unmarshal([]byte(msg.Body), &f)
			b.Bus.Publish("backup", f)

		case "restore":
			//b.Bus.Publish("restore", jmsg.Data)
		case "New.Client":
			fmt.Println("Recieved response")

			err = json.Unmarshal([]byte(msg.Body), &s)
			b.Bus.Publish("newclient", s)
		case "Directory.Scan":
			//var d entity.Directory
			err = json.Unmarshal([]byte(msg.Body), &d)
			b.Bus.Publish("directoryscan", d)
		case "StorageNode.Job":
			fmt.Println("Received client job")
			err = json.Unmarshal([]byte(msg.Body), &s)
			b.Bus.Publish("StorageNodeJob", s)
		case "Client.File":
			//fmt.Println("Received client message")
			err = json.Unmarshal([]byte(msg.Body), &file)
			holder.File = &file
			holder.Type = "clientfile"
			//fmt.Println("Consumed", file)
			b.Bus.Publish("file", holder)
		case "StorageNode.File":
			err = json.Unmarshal([]byte(msg.Body), &file)
			holder.File = &file
			holder.Type = "storagenodefile"
			b.Bus.Publish("file", holder)
			//fmt.Println("Storage node file")
		}

		//fmt.Println("msg consumed", s)
	}
	log.Println("Exiting")
	return nil
}

// func ProducerHandler(chn chan Message, r *Consumer) {
// 	for msgs := range chn {
// 		switch msgs.Type {
// 		case "backup":
// 			r.Bus.Publish(msgs.Type, msgs.Data)

// 		case "restore":
// 			r.Bus.Publish(msgs.Type, msgs.Data)
// 		}
// 	}

// }
