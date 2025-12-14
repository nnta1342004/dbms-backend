package itembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
	itemmodel "hareta/modules/item/model"
)

type createStore interface {
	Create(ctx context.Context, data *itemmodel.Item) error
}
type groupStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*groupitemmodel.GroupItem, error)
}
type createBiz struct {
	logger     logger.Logger
	store      createStore
	groupStore groupStore
}

func NewCreateBiz(store createStore, groupStore groupStore) *createBiz {
	return &createBiz{store: store, groupStore: groupStore, logger: logger.GetCurrent().GetLogger("ItemCreateBiz")}
}

func (biz *createBiz) Create(ctx context.Context, data *itemmodel.ItemCreate) (*itemmodel.Item, error) {
	id, err := appCommon.FromBase58(data.GroupId)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	group, err := biz.groupStore.FindDataWithCondition(ctx, map[string]interface{}{
		"id": id.GetLocalID(),
	})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(groupitemmodel.EntityName, err)
	}

	group.Mask(false)

	item := &itemmodel.Item{
		SQLModel: appCommon.SQLModel{},
		ItemAttr: itemmodel.ItemAttr{
			Name:          data.Name,
			ProductLine:   data.ProductLine,
			Category:      data.Category,
			Quantity:      data.Quantity,
			Discount:      data.Discount,
			Collection:    data.Collection,
			Type:          data.Type,
			Price:         data.Price,
			Color:         data.Color,
			OriginalPrice: data.OriginalPrice,
		},
		ItemSystem: itemmodel.ItemSystem{
			Sold:       0,
			LikeCount:  0,
			CronStatus: 0,
			Tag:        0,
		},
		Description: data.Description,
		AvatarId:    0,
		GroupId:     group.Id,
	}
	if err := biz.store.Create(ctx, item); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(itemmodel.EntityName, err)
	}
	item.Mask(false)
	return item, nil
}
