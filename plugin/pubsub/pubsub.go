package pubsub

import (
	"context"
	"fmt"
	"time"
)

type Topic string

type PubSub interface {
	Publish(ctx context.Context, channel string, data *Message) error
	Subscribe(ctx context.Context, channel string) (ch <-chan *Message, close func() error)
	Reconnect() error
}

type Message struct {
	id        string
	channel   string
	data      map[string]interface{}
	createdAt time.Time
	ackFunc   func() error
	nackFunc  func() error
}

func NewMessage(data map[string]interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (evt *Message) SetChannel(channel string) {
	evt.channel = channel
}

func (evt *Message) String() string {
	return fmt.Sprintf("Message %s value %v", evt.channel, evt.data)
}

func (evt *Message) Channel() string { return evt.channel }

func (evt *Message) Data() map[string]interface{} {
	return evt.data
}

func (evt *Message) SetAckFunc(f func() error) {
	evt.ackFunc = f
}

func (evt *Message) Ack() error {
	return evt.ackFunc()
}

func (evt *Message) SetNackFunc(f func() error) {
	evt.nackFunc = f
}

func (evt *Message) Nack() error {
	return evt.nackFunc()
}
