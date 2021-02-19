package main

import (
	"fmt"

	comms "github.com/wade-sam/fyp-backup-server/externalConsumer"
	client "github.com/wade-sam/fyp-backup-server/pkg/Client"
	"github.com/wade-sam/fyp-backup-server/pkg/Entities"
	"github.com/wade-sam/fyp-backup-server/pkg/policy"
	"github.com/wade-sam/fyp-backup-server/pkg/repository/mongo"
	"github.com/wade-sam/fyp-backup-server/pkg/repository/rabbitmq"
)

func main() {
	servicePersistentRepo, err := mongo.NewMongoRepo("mongodb", "database:27017", "maindb", "root", "fypproject", 10)
	if err != nil {
		fmt.Println(err, "Don't work")
	}

	serviceRabbitRepo, connection, err := rabbitmq.NewRabbitRepo("amqp", "admin", "85v!AP", "rabbitmq", "5672", "/", "host")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(serviceRabbitRepo)
	clientService := client.NewClientService(servicePersistentRepo, serviceRabbitRepo)
	policyService := policy.NewPolicyService(servicePersistentRepo)
	comms.Initialise(clientService, connection)
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
	clientService.UpdateClient("jack", &clientstruct)
	//policyService.CreatePolicy(&policystruct)
	//output, err := clientService.FindClient("sam's macbook pro")
	//fmt.Println(output)
	//clientService.CreateClient(&clientstruct)
}
