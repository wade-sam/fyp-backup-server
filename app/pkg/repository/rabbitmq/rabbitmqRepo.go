package rabbitmq

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/wade-sam/fyp-backup-server/pkg/ports"
)

type BrokerConfig struct {
	Schema         string
	Username       string
	Password       string
	Host           string
	Port           string
	VHost          string
	ConnectionName string
}
type Broker struct {
	config     BrokerConfig
	connection *amqp.Connection
}

func newRabbitClient(config *BrokerConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("%s://%s:%s@%s:%s%s",
		config.Schema,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.VHost,
		//config.ConnectionName,
	))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func NewRabbitRepo(schema, username, password, host, port, vhost, connectioname string) (ports.BrokerRepository, *amqp.Connection, error) {
	configuration := BrokerConfig{
		Schema:         schema,
		Username:       username,
		Password:       password,
		Host:           host,
		Port:           port,
		VHost:          vhost,
		ConnectionName: connectioname,
	}

	repo := &Broker{
		config: configuration,
	}
	client, err := newRabbitClient(&configuration)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Repository.NewRabbitRepo")
	}

	repo.connection = client

	return repo, repo.connection, nil

}

func (r *Broker) GrabInstance() amqp.Connection {

	return *r.connection
}

func (r *Broker) ConnectClient(channel chan int) error {
	return nil
}
