package rabbitmq

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/leductoan3082004/go-sdk/logger"
	"github.com/streadway/amqp"
	"hareta/components/asyncjob"
	"hareta/plugin/pubsub"
)

const RetryLimit = 10

type rabbitMQ struct {
	name       string
	conn       *amqp.Connection
	logger     logger.Logger
	url        string
	initQueues []string
	channel    *amqp.Channel
}

func NewRabbitMQ(name string, q ...string) *rabbitMQ {
	return &rabbitMQ{
		name:       name,
		initQueues: q,
	}
}

func (r *rabbitMQ) GetPrefix() string {
	return r.name
}

func (r *rabbitMQ) Get() interface{} {
	return r
}

func (r *rabbitMQ) Name() string {
	return r.name
}

func (r *rabbitMQ) InitFlags() {
	flag.StringVar(
		&r.url, r.name+"-url", "amqps://zzxkuiqx:NsgCCzZk3zz3J_0gqvxkNAATnHrJxnQ5@armadillo.rmq.cloudamqp.com/zzxkuiqx",
		"RabbitMQ URL",
	)
}

func (r *rabbitMQ) Configure() error {
	r.logger = logger.GetCurrent().GetLogger(r.name)

	r.logger.Infoln("Connecting to RabbitMQ service...")
	conn, err := amqp.Dial(r.url)
	if err != nil {
		r.logger.WithSrc().Fatalln(err)
	}

	r.logger.WithSrc().Infoln("Connected to RabbitMQ service.")

	r.channel, err = conn.Channel()
	if err != nil {
		r.logger.WithSrc().Fatalln(err)
	}

	r.logger.WithSrc().Infoln("Starting to initialize queues...")
	for _, q := range r.initQueues {
		_, err = r.channel.QueueDeclare(
			q,     // queue name
			true,  // durable
			false, // auto delete
			false, // exclusive
			false, // no wait
			nil,   // arguments
		)

		if err != nil {
			r.logger.WithSrc().Fatalln(err)
		}
	}
	r.conn = conn
	return nil
}

func (r *rabbitMQ) reconnect(ctx context.Context) error {
	r.logger.Infoln("Reconnecting to RabbitMQ service...")
	conn, err := amqp.Dial(r.url)

	if err != nil {
		r.logger.WithSrc().Fatalln(err)
	}
	r.channel, err = conn.Channel()
	if err != nil {
		r.logger.WithSrc().Fatalln(err)
	}
	r.logger.WithSrc().Infoln("Reconnected to RabbitMQ service.")
	r.conn = conn
	return nil
}

func (r *rabbitMQ) Reconnect() error {
	job := asyncjob.NewJob(r.reconnect)
	if err := job.ExecuteWithRetry(context.Background()); err != nil {
		r.logger.WithSrc().Fatalln(err) // fatalln here and restart service
	}
	return nil
}

func (r *rabbitMQ) Run() error {
	return r.Configure()
}

func (r *rabbitMQ) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		r.conn.Close()
		r.channel.Close()
	}()
	close(c)
	return c
}

func (r *rabbitMQ) getChannel(ctx context.Context) (*amqp.Channel, error) {
	pubChannel, err := r.conn.Channel()
	if err == amqp.ErrClosed {
		r.logger.Errorln("Channel closed")
		r.Reconnect()
	} else if err != nil {
		r.logger.Errorln(err)
		return nil, err
	} else {
		return pubChannel, nil
	}

	pubChannel, err = r.conn.Channel()
	if err != nil {
		r.logger.Errorln(err)
		return nil, err
	}
	return pubChannel, nil
}

func (r *rabbitMQ) Publish(ctx context.Context, channel string, data *pubsub.Message) error {
	pubChannel, err := r.getChannel(ctx)
	if err != nil {
		r.logger.Errorln(err)
		return err
	}

	msgData, err := json.Marshal(data.Data())

	if err != nil {
		r.logger.Errorln(err)
		return err
	}

	if err := pubChannel.Publish(
		"",      // exchange
		channel, // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msgData,
		},
	); err != nil {
		r.logger.Errorln(err)
		return err
	}

	pubChannel.Close()

	return nil
}

func (r *rabbitMQ) Subscribe(ctx context.Context, channelName string) (
	ch <-chan *pubsub.Message, closeFn func() error,
) {
	msgChan := make(chan *pubsub.Message)
	closeChan := make(chan *amqp.Error)
	r.channel.NotifyClose(closeChan)

	go func() {
		for {
			sub, err := r.channel.Consume(
				channelName, // queue
				"",          // consumer
				false,       // auto-ack
				false,       // exclusive
				false,       // no-local (deprecated, usually false)
				false,       // no-wait
				nil,         // arguments
			)
			if err != nil {
				if err == amqp.ErrClosed {
					r.logger.WithSrc().Errorln("Channel closed, attempting to reconnect...")
					if reconnectErr := r.Reconnect(); reconnectErr != nil {
						r.logger.WithSrc().Errorln("Failed to reconnect:", reconnectErr)
						break
					}
					continue
				}
				r.logger.WithSrc().Errorln("Failed to start consuming:", err)
				continue
			}
			r.channel.NotifyClose(closeChan)

			for {
				select {
				case d := <-sub:
					msgData := make(map[string]interface{})
					if err := json.Unmarshal(d.Body, &msgData); err != nil {
						r.logger.WithSrc().Errorln("Error unmarshalling message:", err)
						if nackErr := d.Nack(false, false); nackErr != nil {
							r.logger.WithSrc().Errorln("Failed to Nack message:", nackErr)
						}
						continue
					}

					appMsg := pubsub.NewMessage(msgData)
					appMsg.SetChannel(channelName)
					appMsg.SetAckFunc(
						func() error {
							return d.Ack(false)
						},
					)
					appMsg.SetNackFunc(
						func() error {
							return d.Nack(false, true)
						},
					)
					msgChan <- appMsg
				case <-ctx.Done():
					return
				case <-closeChan:
					break
				}
			}

		}
	}()

	return msgChan, r.conn.Close
}
