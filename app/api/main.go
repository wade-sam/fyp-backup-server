package main

import (
	"context"
	"fmt"
	"log"
	"time"

	mg "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/mongo"
	rb "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/rabbit"
	"github.com/wade-sam/fyp-backup-server/rabbitBus"
	"github.com/wade-sam/fyp-backup-server/usecase/client"
	cs "github.com/wade-sam/fyp-backup-server/usecase/client"
	"github.com/wade-sam/fyp-backup-server/usecase/policy"

	bs "github.com/wade-sam/fyp-backup-server/usecase/backup"
	//ds "github.com/wade-sam/fyp-backup-server/usecase/dispatcher"
	ps "github.com/wade-sam/fyp-backup-server/usecase/policy"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	creds := options.Credential{
		Username: "root",
		Password: "fypproject",
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s", "mongodb", "database:27017")).SetAuth(creds))
	if err != nil {
		log.Println(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
	}

	configuration := rb.BrokerConfig{
		Schema:         "amqp",
		Username:       "admin",
		Password:       "85v!AP",
		Host:           "rabbitmq",
		Port:           "5672",
		VHost:          "/",
		ConnectionName: "backupserver",
	}

	consumerConf := rb.ConsumerConfig{
		ExchangeName: "main",
		ExchangeType: "direct",
		RoutingKey:   "backupserver",
		QueueName:    "backupserver",
		ConsumerName: "backupserver",
		MaxAttempt:   60,
		Interval:     1 * time.Second,
	}
	producerConf := rb.ProducerConfig{
		ExchangeName: "main",
		ExchangeType: "direct",
		MaxAttempt:   60,
		Interval:     1 * time.Second,
	}

	events := make(map[string]rabbitBus.EventChannelSlice)
	bus := rabbitBus.NewRabbitBus(events)

	ClientsRepo := mg.NewClientMongo(client, "maindb", "clients_collection", 10)
	PolicyMongo := mg.NewPolicyMongo(client, "maindb", "policy_collection", 10)
	BackupMongo := mg.NewBackupMongo(client, "maindb", "backup_collection", 10)
	broker := rb.NewBroker(configuration, producerConf, consumerConf, bus)

	err = broker.Connect()
	if err != nil {
		log.Fatal(err)
	}
	consumer_chan, err := broker.Start()
	go broker.Consume(consumer_chan)

	clientService := cs.NewService(ClientsRepo)
	policyServce := ps.NewService(PolicyMongo)
	policy, err := setupClientPolicy(policyServce, clientService)
	//	policy, err := setupIncClientPolicy(policyServce, clientService)
	//fmt.Println(policy)
	// policy, err := policyServce.GetPolicy("603d7a904b9ac22b561fdcbb")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(policy)

	backup := bs.NewService(ClientsRepo, PolicyMongo, BackupMongo, broker, bus)
	err = backup.StartBackup(policy, "Full")
	if err != nil {
		fmt.Println(err)
	}
	// dispatcherService := ds.NewService(broker, bus)
	// // fmt.Println("starting dispatch")
	// clientname, err := dispatcherService.SearchForNewClient()
	// scanresult, err := dispatcherService.GetDirectoryScan(clientname)
	// fmt.Println("client name is: ", scanresult)
}

func setupClientPolicy(p *policy.Service, c *client.Service) (string, error) {
	client, err := c.CreateClient("samwade", "newclient")
	if err != nil {
		return "", err
	}
	fullbackup := []string{"Monday", "Friday"}
	clients := []string{client}
	policy, err := p.CreatePolicy("Friday Backup", "full", 10, fullbackup, []string{}, clients)
	if err != nil {
		return "", err
	}
	rclient, _ := c.GetClient(client)
	rclient.AddPolicy(policy)
	c.UpdateClient(rclient)
	return policy, nil
}

func setupIncClientPolicy(p *policy.Service, c *client.Service) (string, error) {
	client, err := c.CreateClient("samwade", "newclient")
	if err != nil {
		return "", err
	}
	IncBackup := []string{"Monday", "Friday"}
	FullBackup := []string{"Sunday", "Tuesday"}
	clients := []string{client}
	policy, err := p.CreatePolicy("Friday Backup", "both", 10, FullBackup, IncBackup, clients)
	if err != nil {
		return "", err
	}
	rclient, _ := c.GetClient(client)
	rclient.AddPolicy(policy)
	c.UpdateClient(rclient)
	return policy, nil
}
