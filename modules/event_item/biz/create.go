package eventitembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
	eventitemmodel "hareta/modules/event_item/model"
	itemmodel "hareta/modules/item/model"
)

type createStore interface {
	Create(ctx context.Context, data *eventitemmodel.EventItem) error
}
type createItemStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*itemmodel.Item, error)
}
type createEventStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*eventmodel.Event, error)
}
type createBiz struct {
	store      createStore
	itemStore  createItemStore
	eventStore createEventStore
	logger     logger.Logger
}

func NewCreateBiz(store createStore, eventStore createEventStore, itemStore createItemStore) *createBiz {
	return &createBiz{
		store:      store,
		logger:     logger.GetCurrent().GetLogger("EventItemCreateBiz"),
		eventStore: eventStore,
		itemStore:  itemStore,
	}
}

func (biz *createBiz) Create(ctx context.Context, data *eventitemmodel.EventItemCreate) error {
	eventId, err := appCommon.FromBase58(data.EventId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	itemId, err := appCommon.FromBase58(data.ItemId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	// Check event exist
	_, err = biz.eventStore.FindDataWithCondition(ctx, map[string]interface{}{
		"id": eventId.GetLocalID(),
	})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(eventmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(eventmodel.EntityName, err)
	}

	// Check item exist
	_, err = biz.itemStore.FindDataWithCondition(ctx, map[string]interface{}{
		"id": itemId.GetLocalID(),
	})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(itemmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(itemmodel.EntityName, err)
	}

	createData := &eventitemmodel.EventItem{
		EventId: int64(eventId.GetLocalID()),
		ItemId:  int64(itemId.GetLocalID()),
	}
	if err := biz.store.Create(ctx, createData); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotCreateEntity(eventitemmodel.EntityName, err)
	}
	return nil
}
