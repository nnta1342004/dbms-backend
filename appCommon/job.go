package appCommon

import (
	"context"
	"hareta/plugin/pubsub"
)

type SubJob struct {
	Title     string
	NumWorker int
	Topic     string
	Hld       func(ctx context.Context, message *pubsub.Message, workerId int) error
}
