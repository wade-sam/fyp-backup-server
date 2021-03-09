package rabbit

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/streadway/amqp"
	"github.com/wade-sam/fyp-backup-server/entity"
)

type ProducerConfig struct {
	ExchangeName string
	ExchangeType string
	MaxAttempt   int
	Interval     time.Duration
	connection   *amqp.Connection
}

func (b *Broker) Publish(routekey, messagetype string, body *DTO) error {
	channel, err := b.Channel()
	if err != nil {
		return err
	}
	dto, err := Serialize(body)
	if err != nil {
		return err
	}
	defer channel.Close()
	err = channel.Publish(
		b.Producer.ExchangeName,
		routekey,
		false,
		false,
		amqp.Publishing{
			Type:        messagetype,
			ContentType: "application/json",
			Body:        dto,
		})
	if err != nil {
		return err
	}
	return nil

}

func (b *Broker) SearchForNewClient() (string, error) {
	chn, err := b.Bus.Subscribe("newclient")
	if err != nil {
		return "", entity.ErrNoMatchingTopic
	}
	messagetype := "New.Client"
	key := "newclient"
	dto := DTO{}
	err = b.Publish(key, messagetype, &dto)

	if err != nil {
		return "", err
	}

	for i := 1; i < 10; i++ {
		select {
		case msg := <-chn:
			s := ""
			mapstructure.Decode(msg.Data, &s)
			fmt.Println("mastructure", s)
			return s, nil
		default:
			time.Sleep(2 * time.Second)
		}
	}
	close(chn)
	fmt.Println("NO NEW CLIENT")
	return "", entity.ErrNoNewClient
}

func (b *Broker) DirectoryScan(client string) (*entity.Directory, error) {
	chn, err := b.Bus.Subscribe("directoryscan")
	if err != nil {
		return nil, entity.ErrNoMatchingTopic
	}
	messagetype := "Directory.Scan"
	key := fmt.Sprintf("%v", client)
	dto := DTO{}
	err = b.Publish(key, messagetype, &dto)
	if err != nil {
		return nil, err
	}
	for i := 1; i < 200; i++ {
		select {
		case msg := <-chn:
			d := entity.Directory{}
			mapstructure.Decode(msg.Data, &d)
			return &d, nil
			//scanresult, err := msg.Data.(*entity.Directory)
			// if err != nil {
			// 	fmt.Println(scanresult)
			// 	return scanresult, nil
			// }

		default:
			time.Sleep(2 * time.Second)
		}
	}
	close(chn)
	return nil, entity.ErrNotFound

}
