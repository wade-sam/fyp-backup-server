package rabbitBus

import (
	"sync"

	"github.com/wade-sam/fyp-backup-server/entity"
)

type Event struct {
	Data  interface{}
	Topic string
}

type EventChannel chan Event
type EventChannelSlice []EventChannel

type RabbitBus struct {
	subscribers map[string]EventChannelSlice
	rm          sync.RWMutex
}

func NewRabbitBus(subscribers map[string]EventChannelSlice) *RabbitBus {
	return &RabbitBus{
		subscribers: subscribers,
	}
}

func publish_send(data Event, eventChannelSlices EventChannelSlice) {
	for _, ch := range eventChannelSlices {
		ch <- data
	}
}

func (rb *RabbitBus) Publish(topic string, data interface{}) error {
	rb.rm.Lock()
	if chans, found := rb.subscribers[topic]; found {
		channels := append(EventChannelSlice{}, chans...)
		if len(channels) > 0 {
			go publish_send(Event{Data: data, Topic: topic}, channels)
			rb.rm.Unlock()
			return nil
		}
		return entity.ErrNoSubscribersForTopic
	}
	rb.rm.Unlock()
	return entity.ErrNoMatchingTopic
}

func (rb *RabbitBus) Subscribe(topic string) (EventChannel, error) {
	ch := make(chan Event)
	rb.rm.Lock()
	if prev, found := rb.subscribers[topic]; found {
		rb.subscribers[topic] = append(prev, ch)
	} else {
		rb.subscribers[topic] = append([]EventChannel{}, ch)
	}
	rb.rm.Unlock()

	return ch, nil
}

func (rb *RabbitBus) Unsubscribe(topic string, ch chan Event) error {
	if _, found := rb.subscribers[topic]; found {
		for i := range rb.subscribers[topic] {
			if rb.subscribers[topic][i] == ch {
				rb.subscribers[topic] = append(rb.subscribers[topic][:i], rb.subscribers[topic][i+1:]...)
				return nil
			}
		}
		return entity.ErrNotFound
	}
	return entity.ErrNoMatchingTopic
}
