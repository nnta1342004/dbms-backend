package itembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

type makeDefaultStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*itemmodel.Item, error)
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}

type makeDefaultBiz struct {
	store  makeDefaultStore
	logger logger.Logger
}

func NewMakeDefaultBiz(store makeDefaultStore) *makeDefaultBiz {
	return &makeDefaultBiz{logger: logger.GetCurrent().GetLogger("MakeDefaultItemBiz"), store: store}
}

func (biz *makeDefaultBiz) MakeDefault(ctx context.Context, data *itemmodel.ItemMakeDefault) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	item, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"id": id.GetLocalID(),
	})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(itemmodel.EntityName, err)
	}

	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{
		"group_id": item.GroupId,
	}, map[string]interface{}{
		"default": 0,
	}); err != nil {
		return appCommon.ErrCannotUpdateEntity(itemmodel.EntityName, err)
	}

	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{
		"id": item.Id,
	}, map[string]interface{}{
		"default": 1,
	}); err != nil {
		return appCommon.ErrCannotUpdateEntity(itemmodel.EntityName, err)
	}
	return nil
}
