package eventbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
	"time"
)

type createStore interface {
	Create(ctx context.Context, data *eventmodel.Event) error
}
type createBiz struct {
	store  createStore
	logger logger.Logger
}

func NewCreateBiz(store createStore) *createBiz {
	return &createBiz{store: store, logger: logger.GetCurrent().GetLogger("CreateEventBiz")}
}

func (biz *createBiz) Create(ctx context.Context, data *eventmodel.EventCreate) (*eventmodel.Event, error) {

	dateStart := time.Unix(data.DateStart, 0)
	dateEnd := time.Unix(data.DateEnd, 0)
	event := &eventmodel.Event{
		DateStart:      dateStart,
		DateEnd:        dateEnd,
		OverallContent: data.OverallContent,
		DetailContent:  data.DetailContent,
		Discount:       data.Discount,
		Avatar:         data.Avatar,
	}

	if err := biz.store.Create(ctx, event); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(eventmodel.EntityName, err)
	}

	event.Mask(false)
	return event, nil
}
