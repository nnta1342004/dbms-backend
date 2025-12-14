package eventitembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	eventitemmodel "hareta/modules/event_item/model"
)

type deleteStore interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
}
type deleteBiz struct {
	store  deleteStore
	logger logger.Logger
}

func NewDeleteBiz(store deleteStore) *deleteBiz {
	return &deleteBiz{store: store, logger: logger.GetCurrent().GetLogger("EventItemDeleteBiz")}
}

func (biz *deleteBiz) Delete(ctx context.Context, data *eventitemmodel.EventItemDelete) error {
	eventId, err := appCommon.FromBase58(data.EventId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	itemId, err := appCommon.FromBase58(data.ItemId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	conditions := map[string]interface{}{
		"event_id": eventId.GetLocalID(),
		"item_id":  itemId.GetLocalID(),
	}

	if err := biz.store.DeleteWithCondition(ctx, conditions); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(eventitemmodel.EntityName, err)
	}
	return nil
}
