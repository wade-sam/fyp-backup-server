package rabbitBus_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
	repo "github.com/wade-sam/fyp-backup-server/rabbitBus"
)

func Test_Subscribe(t *testing.T) {
	subs := map[string]repo.EventChannelSlice{}
	rb := repo.NewRabbitBus(subs)
	_, err := rb.Subscribe("backups")
	assert.Nil(t, err)
}

func Test_Publish(t *testing.T) {
	subs := map[string]repo.EventChannelSlice{}
	rb := repo.NewRabbitBus(subs)
	_, err := rb.Subscribe("backups")
	assert.Nil(t, err)
	err = rb.Publish("backups", "policy1")
	assert.Nil(t, err)
	err = rb.Publish("policies", "policy1")
	assert.Equal(t, entity.ErrNoMatchingTopic, err)
}

func Test_Consuming(t *testing.T) {
	msg := "msg1"
	subs := map[string]repo.EventChannelSlice{}
	rb := repo.NewRabbitBus(subs)
	ch, err := rb.Subscribe("backups")
	ch2, err := rb.Subscribe("backups")
	assert.Nil(t, err)
	go test_publish(rb, "backups", msg)
	for i := 1; i < 2; i++ {
		d := <-ch
		assert.Equal(t, msg, d.Data)
		fmt.Println(d.Data)
		e := <-ch2
		assert.Equal(t, msg, e.Data)
		fmt.Println(e.Data)
	}
}

func Test_Unsubscribe(t *testing.T) {
	subs := map[string]repo.EventChannelSlice{}
	rb := repo.NewRabbitBus(subs)
	ch, err := rb.Subscribe("backups")
	assert.Nil(t, err)
	err = rb.Unsubscribe("backups", ch)
	assert.Nil(t, err)
	err = rb.Unsubscribe("backups", ch)
	assert.Equal(t, entity.ErrNotFound, err)
	err = rb.Unsubscribe("testing", ch)
	assert.Equal(t, entity.ErrNoMatchingTopic, err)
	err = rb.Publish("backups", "hello")
	assert.Equal(t, entity.ErrNoSubscribersForTopic, err)
}

func test_publish(rb *repo.RabbitBus, topic string, data string) {
	time.Sleep(time.Duration(1) * time.Second)
	rb.Publish(topic, data)
}
