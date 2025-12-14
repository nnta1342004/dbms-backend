package eventbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
	"time"
)

type updateStore interface {
	UpdateEvent(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}
type updateBiz struct {
	store  updateStore
	logger logger.Logger
}

func NewUpdateBiz(store updateStore) *updateBiz {
	return &updateBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("EventUpdateBiz"),
	}
}

func (biz *updateBiz) Update(ctx context.Context, data *eventmodel.EventUpdate) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	updateData := make(map[string]interface{})
	if data.DetailContent != nil {
		updateData["detail_content"] = *data.DetailContent
	}
	if data.OverallContent != nil {
		updateData["overall_content"] = *data.OverallContent
	}
	if data.DateStart != nil {
		updateData["date_start"] = time.Unix(*data.DateStart, 0)
	}
	if data.DateEnd != nil {
		updateData["date_end"] = time.Unix(*data.DateEnd, 0)
	}
	if data.Discount != nil {
		updateData["discount"] = *data.Discount
	}
	if data.Avatar != nil {
		updateData["avatar"] = *data.Avatar
	}
	conditions := map[string]interface{}{
		"id": id.GetLocalID(),
	}
	if err := biz.store.UpdateEvent(ctx, conditions, updateData); err != nil {
		biz.logger.Errorln(err)
		return appCommon.ErrCannotUpdateEntity(eventmodel.EntityName, err)
	}
	return nil
}
