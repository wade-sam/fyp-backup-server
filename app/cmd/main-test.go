package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	comms "github.com/wade-sam/fyp-backup-server/ExternalConsumer"
	client "github.com/wade-sam/fyp-backup-server/pkg/Client"
	"github.com/wade-sam/fyp-backup-server/pkg/Entities"
	"github.com/wade-sam/fyp-backup-server/pkg/policy"
	"github.com/wade-sam/fyp-backup-server/pkg/repository/mongo"
	"github.com/wade-sam/fyp-backup-server/pkg/repository/rabbitmq"
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
func main() {
	servicePersistentRepo, err := mongo.NewMongoRepo("mongodb", "database:27017", "maindb", "root", "fypproject", 10)
	if err != nil {
		fmt.Println(err, "Don't work")
	}

	configuration := BrokerConfig{
		Schema:         "amqp",
		Username:       "admin",
		Password:       "85v!AP",
		Host:           "rabbitmq",
		Port:           "5672",
		VHost:          "/",
		ConnectionName: "backupserver",
	}
	conn, err := newRabbitClient(&configuration)
	//serviceRabbitRepo, connection, err := rabbitmq.NewRabbitRepo("amqp", "admin", "85v!AP", "rabbitmq", "5672", "/", "host")
	if err != nil {
		log.Println(err, "Main.NewRabbitConnection")
	}
	serviceRabbitRepo := rabbitmq.NewRabbitRepo(conn)

	fmt.Println(serviceRabbitRepo)
	IclientService, EclientService := client.NewClientService(servicePersistentRepo, serviceRabbitRepo)
	policyService := policy.NewPolicyService(servicePersistentRepo)
	//fmt.Println(connection)
	go comms.Initialise(EclientService, conn)
	clientstruct := Entities.Client{
		Clientname:   "jackie boy",
		Consumername: "host1",
	}

	policystruct := Entities.Policy{
		Policyname:  "Wednesday Backup",
		Clients:     []string{"sam macbook pro", "cameron's macbook pro", "pippa's macbook pro"},
		Retention:   200,
		Scale:       "monthly",
		Fullbackup:  []string{"Monday", "Friday"},
		Incremental: []string{"Tuesday", "Wednesday", "Thursday", "Saturday", "Sunday"},
	}
	policyService.CreatePolicy(&policystruct)
	policyService.UpdatePolicy("Wednesday Backup", &policystruct)
	//policyService.DeletePolicy("Wednesday Backup")
	IclientService.UpdateClient("jack", &clientstruct)
	//policyService.CreatePolicy(&policystruct)
	//output, err := clientService.FindClient("sam's macbook pro")
	//fmt.Println(output)
	//clientService.CreateClient(&clientstruct)

	select {}
}
