package main

import (
	"context"
	"fmt"
	"log"

	"net/http"
	"time"

	mg "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/mongo"
	rb "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/rabbit"

	"github.com/wade-sam/fyp-backup-server/api/handler"
	"github.com/wade-sam/fyp-backup-server/rabbitBus"
	cs "github.com/wade-sam/fyp-backup-server/usecase/client"

	bs "github.com/wade-sam/fyp-backup-server/usecase/backup"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	ds "github.com/wade-sam/fyp-backup-server/usecase/dispatcher"
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

	//Initialise Services
	clientService := cs.NewService(ClientsRepo)
	policyService := ps.NewService(PolicyMongo)
	dispatchService := ds.NewService(broker, bus)
	//policy, err := setupClientPolicy(policyService, clientService)
	backup := bs.NewService(ClientsRepo, PolicyMongo, BackupMongo, broker, bus)

	err = broker.Connect()
	if err != nil {
		log.Fatal(err)
	}
	consumer_chan, err := broker.Start()
	go broker.Consume(consumer_chan)

	//Initialise router
	router := mux.NewRouter().StrictSlash(true)
	//router.Use(CORS)
	handler.MakeClientHandlers(router, clientService, policyService, dispatchService)
	handler.MakePolicyHolders(router, clientService, policyService)
	handler.MakeBackupHandlers(router, clientService, policyService, backup)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:3000"}))(router)))

	// policy, err := setupIncClientPolicy(policyService, clientService)
	// fmt.Println(policy)
	//returnPolicy, err := policyService.GetPolicy(policy)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(returnPolicy)

	// err = backup.StartBackup("607ed003083610a78abad050", "Full")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// dispatcherService := ds.NewService(broker, bus)
	// // fmt.Println("starting dispatch")
	// clientname, err := dispatcherService.SearchForNewClient()
	// scanresult, err := dispatcherService.GetDirectoryScan(clientname)
	// fmt.Println("client name is: ", scanresult)
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		// w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		// w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			log.Println("REACHED")
			// w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
			// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			// w.Header().Set("Content-Type", "application/json")
			//w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		//return
	})
}

// func setupClientPolicy(p *policy.Service, c *client.Service) (string, error) {
// 	client, err := c.CreateClient("samwade", "newclient")
// 	if err != nil {
// 		return "", err
// 	}
// 	fullbackup := []string{"Monday", "Friday"}
// 	clients := []string{client}
// 	policy, err := p.CreatePolicy("Friday Backup", "full", 10, fullbackup, []string{}, clients)
// 	if err != nil {
// 		return "", err
// 	}
// 	rclient, _ := c.GetClient(client)
// 	rclient.AddPolicy(policy)
// 	c.UpdateClient(rclient)
// 	return policy, nil
// }

// func setupIncClientPolicy(p *policy.Service, c *client.Service) (string, error) {
// 	client, err := c.CreateClient("samwade", "newclient")
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println("Created client", client)
// 	IncBackup := []string{"Monday", "Friday"}
// 	FullBackup := []string{"Sunday", "Tuesday"}
// 	clients := []string{client}
// 	policy, err := p.CreatePolicy("Friday Backup", "both", 10, FullBackup, IncBackup, clients)
// 	if err != nil {
// 		log.Println(err)
// 		return "", err
// 	}
// 	rclient, _ := c.GetClient(client)
// 	rclient.AddPolicy(policy)
// 	c.UpdateClient(rclient)
// 	return policy, nil
// }
