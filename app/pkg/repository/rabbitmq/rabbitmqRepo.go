package rabbitmq

import (

	//"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/wade-sam/fyp-backup-server/pkg/ports"
)

type Broker struct {
	connection *amqp.Connection
}

func NewRabbitRepo(connection *amqp.Connection) ports.BrokerRepository {
	repo := &Broker{}
	repo.connection = connection

	return repo

}

func (r *Broker) GrabInstance() amqp.Connection {

	return *r.connection
}

func (r *Broker) ConnectClient(channel chan int) error {
	return nil
}
