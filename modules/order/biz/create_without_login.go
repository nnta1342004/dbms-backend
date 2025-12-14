package orderbiz

import (
	"context"
	"errors"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
	itemordermodel "hareta/modules/item_order/model"
	ordermodel "hareta/modules/order/model"
)

type createWithoutLoginStore interface {
	Create(ctx context.Context, data *ordermodel.Order) error
}
type createWithoutLoginItemStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
	ListInIds(ctx context.Context, ids []int64, conditions map[string]interface{}, moreInfo ...string) ([]itemmodel.Item, error)
}
type createItemOrderWithoutLoginStore interface {
	CreateMany(ctx context.Context, data []itemordermodel.ItemOrder) error
}
type createWithoutLoginBiz struct {
	store          createWithoutLoginStore
	itemStore      createWithoutLoginItemStore
	itemOrderStore createItemOrderWithoutLoginStore
	logger         logger.Logger
}

func NewCreateWithoutLoginBiz(store createWithoutLoginStore, itemStore createWithoutLoginItemStore, itemOrderStore createItemOrderWithoutLoginStore) *createWithoutLoginBiz {
	return &createWithoutLoginBiz{
		store:          store,
		itemStore:      itemStore,
		itemOrderStore: itemOrderStore,
		logger:         logger.GetCurrent().GetLogger("CreateOrderWithoutLoginBiz"),
	}
}

func (biz *createWithoutLoginBiz) Create(ctx context.Context, data *ordermodel.OrderCreateWithoutLogin) (*ordermodel.Order, error) {
	if len(data.Item) == 0 {
		return nil, appCommon.ErrInvalidRequest(errors.New("length of items must be greater than zero"))
	}
	itemIds := make([]int64, len(data.Item))

	for i, val := range data.Item {
		id, err := appCommon.FromBase58(val.Id)
		if err != nil {
			return nil, appCommon.ErrInvalidRequest(err)
		}
		itemIds[i] = int64(id.GetLocalID())
	}

	items, err := biz.itemStore.ListInIds(ctx, itemIds, nil)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(itemmodel.EntityName, err)
	}

	if len(items) != len(itemIds) {
		return nil, appCommon.ErrInvalidRequest(errors.New("length of item does not match system length"))
	}

	mp := make(map[int64]int64)
	for i, val := range data.Item {
		mp[itemIds[i]] = val.Quantity
	}

	// calculate price
	var total int64
	total = 0
	for i := range items {
		if items[i].Quantity < mp[items[i].Id] {
			return nil, ordermodel.ErrQuantityLimitExceed
		}
		total += mp[items[i].Id] * (items[i].Price - items[i].Discount)
	}

	// create order with that price
	order := &ordermodel.Order{
		Address: data.Address,
		Name:    data.Name,
		Email:   data.Email,
		Phone:   data.Phone,
		Total:   total,
	}

	if err := biz.store.Create(ctx, order); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	// create item order
	itemOrder := make([]itemordermodel.ItemOrder, len(itemIds))
	for i := range itemOrder {
		itemOrder[i] = itemordermodel.ItemOrder{
			ItemId:   itemIds[i],
			OrderId:  order.Id,
			Quantity: mp[itemIds[i]],
		}
	}
	if err := biz.itemOrderStore.CreateMany(ctx, itemOrder); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(itemordermodel.EntityName, err)
	}

	// subtract the quantity to each item respectively
	for i := range items {
		if err := biz.itemStore.UpdateWithCondition(ctx, map[string]interface{}{
			"id": items[i].Id,
		}, map[string]interface{}{
			"quantity": items[i].Quantity - mp[items[i].Id],
		}); err != nil {
			biz.logger.WithSrc().Errorln(err)
			return nil, appCommon.ErrCannotUpdateEntity(itemmodel.EntityName, err)
		}
	}

	order.Mask(false)
	return order, nil
}
