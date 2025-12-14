package eventbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
)

type deleteStore interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
}
type deleteBiz struct {
	store  deleteStore
	logger logger.Logger
}

func NewDeleteBiz(store deleteStore) *deleteBiz {
	return &deleteBiz{store: store, logger: logger.GetCurrent().GetLogger("DeleteEventBiz")}
}

func (biz *deleteBiz) Delete(ctx context.Context, data *eventmodel.EventDelete) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	if err := biz.store.DeleteWithCondition(ctx, map[string]interface{}{"id": id.GetLocalID()}); err != nil {
		biz.logger.Errorln(err)
		return appCommon.ErrCannotDeleteEntity(eventmodel.EntityName, err)
	}
	return nil
}
