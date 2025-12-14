package eventbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
)

type findEventStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*eventmodel.Event, error)
}

type findEventBiz struct {
	store  findEventStore
	logger logger.Logger
}

func NewFindEventBiz(store findEventStore) *findEventBiz {
	return &findEventBiz{store: store}
}

func (biz *findEventBiz) FindDataWithCondition(ctx context.Context, id string) (*eventmodel.Event, error) {
	realId, err := appCommon.FromBase58(id)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	data, err := biz.store.FindDataWithCondition(
		ctx,
		map[string]interface{}{"id": realId.GetLocalID()},
		"Items", "Items.Item", "Items.Item.Avatar",
	)
	if err != nil {
		if err == appCommon.RecordNotFound {
			return nil, appCommon.ErrEntityNotFound(eventmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(eventmodel.EntityName, err)
	}

	data.Mask(false)
	for i := range data.Items {
		data.Items[i].Item.Mask(false)
		if data.Items[i].Item.Avatar != nil {
			data.Items[i].Item.Avatar.Mask(false)
		}
	}
	return data, nil
}
