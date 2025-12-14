package subscriber

import (
	"context"
	"fmt"
	goservice "github.com/leductoan3082004/go-sdk"
	logger2 "github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	"hareta/components/asyncjob"
	"hareta/plugin/pubsub"
	passwordrecoverysub "hareta/subscriber/send_mail_recovey_password"
)

type pbEngine struct {
	serviceCtx goservice.ServiceContext
}

func NewEngine(serviceCtx goservice.ServiceContext) *pbEngine {
	return &pbEngine{serviceCtx: serviceCtx}
}

func (engine *pbEngine) Start() error {
	return engine.startSubWorkerTopic(
		passwordrecoverysub.SendMailRecoveryPassword(engine.serviceCtx, appCommon.TopicSendMailRecoveryPassword, 3),
	)
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *pbEngine) startSubWorkerTopic(jobs ...appCommon.SubJob) error {
	ps := engine.serviceCtx.MustGet(appCommon.PluginRabbitMQ).(pubsub.PubSub)
	logger := logger2.GetCurrent().GetLogger("subscriber-engine")

	getJobHandler := func(job *appCommon.SubJob, message *pubsub.Message, workerId int) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in handler", r)
				}
			}()
			return job.Hld(ctx, message, workerId)
		}
	}

	for i, item := range jobs {
		logger.Infoln("---------------------------")
		for id := 0; id < item.NumWorker; id++ {
			logger.Infoln("✅  Setup subscriber for:", item.Title, "|", "Worker id:", id)

			go func(i, id int) {
				c, _ := ps.Subscribe(context.Background(), jobs[i].Topic)
				for {
					msg := <-c
					jobHdl := getJobHandler(&jobs[i], msg, id)
					job := asyncjob.NewJob(jobHdl)
					if err := job.ExecuteWithRetry(context.Background()); err != nil {
						logger.WithSrc().Errorln(err)
						if err := msg.Nack(); err != nil {
							logger.WithSrc().Errorln(err)
						}
					} else {
						if err := msg.Ack(); err != nil {
							logger.WithSrc().Errorln("error when ACK", err, msg.Data(), msg)
						}
					}
				}
			}(i, id)
		}
		logger.Infoln("---------------------------")
	}
	return nil
}
