package rabbit_test

import (
	"testing"
	"time"

	repo "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/rabbit"
	r "github.com/wade-sam/fyp-backup-server/Infrastructure/Repositories/rabbitBus"
)

func InitialiseRepo() *repo.Broker {
	configuration := repo.BrokerConfig{
		Schema:         "amqp",
		Username:       "admin",
		Password:       "85v!AP",
		Host:           "rabbitmq",
		Port:           "5672",
		VHost:          "/",
		ConnectionName: "backupserver",
	}
	broker := repo.NewBroker(configuration)
	return broker
}

func ConsumerConfig() *repo.ConsumerConfig {
	return &repo.ConsumerConfig{
		ExchangeName: "main",
		ExchangeType: "direct",
		RoutingKey:   "backupserver",
		QueueName:    "backupserver",
		ConsumerName: "backupserver",
		MaxAttempt:   60,
		Interval:     1 * time.Second,
	}
}

func Test_InitialiseConsumer(t *testing.T) {
	subs := map[string]r.EventChannelSlice{}
	rbs := r.NewRabbitBus(subs)
	rb := InitialiseRepo()
	config := ConsumerConfig()
	consumer := repo.NewRabbitConsumer(*config, rb, rbs)
	err := consumer.Start()

}
