package cartbiz

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
	itemmodel "hareta/modules/item/model"
	usermodel "hareta/modules/user/model"
)

type updateStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*cartmodel.Cart, error)
}
type itemUpdateStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*itemmodel.Item, error)
}
type updateBiz struct {
	store     updateStore
	itemStore itemUpdateStore
}

func NewUpdateBiz(store updateStore, itemStore itemUpdateStore) *updateBiz {
	return &updateBiz{store: store, itemStore: itemStore}
}
func (biz *updateBiz) Update(ctx context.Context, user *usermodel.User, data *cartmodel.CartUpdate) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	cart, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id.GetLocalID(), "user_id": user.Id})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(cartmodel.EntityName, err)
		}
		return appCommon.ErrCannotGetEntity(cartmodel.EntityName, err)
	}

	item, err := biz.itemStore.FindDataWithCondition(ctx, map[string]interface{}{"id": cart.ItemId})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(itemmodel.EntityName, err)
		}
		return appCommon.ErrCannotGetEntity(itemmodel.EntityName, err)
	}

	if data.Quantity > item.Quantity {
		return cartmodel.ErrQuantityExceed
	}

	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{"id": id.GetLocalID()}, map[string]interface{}{
		"quantity": data.Quantity,
	}); err != nil {
		return appCommon.ErrCannotUpdateEntity(cartmodel.EntityName, err)
	}
	return nil
}
