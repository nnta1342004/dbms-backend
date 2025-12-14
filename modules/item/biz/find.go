package itembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

type findStore interface {
	FindDataById(ctx context.Context, id int64, moreInfo ...string) (*itemmodel.Item, error)
}

type findBiz struct {
	store  findStore
	logger logger.Logger
}

func NewFindBiz(store findStore) *findBiz {
	return &findBiz{store: store}
}

func (biz *findBiz) Find(ctx context.Context, filter *itemmodel.ItemFind) (*itemmodel.Item, error) {
	id, err := appCommon.FromBase58(filter.Id)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	item, err := biz.store.FindDataById(ctx, int64(id.GetLocalID()), "Avatar", "Group")
	if err != nil {
		if err == appCommon.RecordNotFound {
			return nil, appCommon.ErrEntityNotFound(itemmodel.EntityName, err)
		}
		biz.logger.WithSrc().Error(err)
		return nil, appCommon.ErrCannotGetEntity(itemmodel.EntityName, err)
	}

	item.Mask(false)
	if item.Avatar != nil {
		item.Avatar.Mask(false)
	}
	item.Group.Mask(false)
	item.HideSensitiveData()
	return item, nil
}
