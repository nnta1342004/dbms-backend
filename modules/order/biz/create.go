package orderbiz

import (
	"context"
	"errors"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
	itemmodel "hareta/modules/item/model"
	itemordermodel "hareta/modules/item_order/model"
	ordermodel "hareta/modules/order/model"
	usermodel "hareta/modules/user/model"
)

type createStore interface {
	Create(ctx context.Context, data *ordermodel.Order) error
}
type itemStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
	ListInIds(ctx context.Context, ids []int64, conditions map[string]interface{}, moreInfo ...string) ([]itemmodel.Item, error)
}
type cartListStore interface {
	ListInIds(ctx context.Context, ids []int64, moreInfo ...string) ([]cartmodel.Cart, error)
	DeleteInIds(ctx context.Context, ids []int64) error
}
type itemOrderStore interface {
	CreateMany(ctx context.Context, data []itemordermodel.ItemOrder) error
}
type createBiz struct {
	itemOrderStore itemOrderStore
	store          createStore
	itemStore      itemStore
	cartStore      cartListStore
	logger         logger.Logger
}

func NewCreateBiz(store createStore, cartStore cartListStore, itemStore itemStore, itemOrderStore itemOrderStore) *createBiz {
	return &createBiz{
		itemStore:      itemStore,
		store:          store,
		cartStore:      cartStore,
		itemOrderStore: itemOrderStore,
		logger:         logger.GetCurrent().GetLogger("OrderCreateBiz"),
	}
}

func (biz *createBiz) Create(ctx context.Context, user *usermodel.User, data *ordermodel.OrderCreate) (*ordermodel.Order, error) {
	cartIds := make([]int64, len(data.Id))
	for i := range cartIds {
		id, err := appCommon.FromBase58(data.Id[i])
		if err != nil {
			return nil, appCommon.ErrInvalidRequest(err)
		}
		cartIds[i] = int64(id.GetLocalID())
	}

	cartItems, err := biz.cartStore.ListInIds(ctx, cartIds, "Item")
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(cartmodel.EntityName, err)
	}

	if cartItems == nil || len(cartItems) != len(data.Id) {
		return nil, appCommon.ErrInvalidRequest(errors.New("len of data ids does not match with system ids"))
	}

	for i := range cartItems {
		if cartItems[i].Quantity > cartItems[i].Item.Quantity {
			return nil, cartmodel.ErrQuantityExceed
		}
		if err := biz.itemStore.UpdateWithCondition(ctx, map[string]interface{}{
			"id": cartItems[i].Item.Id,
		}, map[string]interface{}{
			"quantity": cartItems[i].Item.Quantity - cartItems[i].Quantity,
		}); err != nil {
			biz.logger.WithSrc().Errorln(err)
			return nil, appCommon.ErrCannotUpdateEntity(itemmodel.EntityName, err)
		}
	}
	var totalCost int64
	for _, val := range cartItems {
		totalCost += val.Quantity * val.Item.Price
	}

	order := &ordermodel.Order{
		Address: data.Address,
		Name:    data.Name,
		Email:   data.Email,
		Phone:   data.Phone,
		UserId:  user.Id,
		Total:   totalCost,
	}

	if err := biz.cartStore.DeleteInIds(ctx, cartIds); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}
	if err := biz.store.Create(ctx, order); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	itemOrder := make([]itemordermodel.ItemOrder, len(cartItems))
	for i := range itemOrder {
		itemOrder[i] = itemordermodel.ItemOrder{
			ItemId:   cartItems[i].ItemId,
			OrderId:  order.Id,
			Quantity: cartItems[i].Quantity,
		}
	}
	if err := biz.itemOrderStore.CreateMany(ctx, itemOrder); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(itemordermodel.EntityName, err)
	}
	order.Mask(false)
	return order, nil
}
