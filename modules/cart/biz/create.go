package cartbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
	itemmodel "hareta/modules/item/model"
	usermodel "hareta/modules/user/model"
)

type createStore interface {
	Create(ctx context.Context, data *cartmodel.Cart) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*cartmodel.Cart, error)
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}
type itemCreateStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*itemmodel.Item, error)
}
type createBiz struct {
	store     createStore
	itemStore itemCreateStore
	logger    logger.Logger
}

func NewCreateBiz(store createStore, store2 itemCreateStore) *createBiz {
	return &createBiz{store: store, itemStore: store2, logger: logger.GetCurrent().GetLogger("AddItemToCartBiz")}
}
func (biz *createBiz) Create(ctx context.Context, user *usermodel.User, data *cartmodel.CartCreate) (*cartmodel.Cart, error) {
	itemId, err := appCommon.FromBase58(data.ItemId)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	item, err := biz.itemStore.FindDataWithCondition(ctx, map[string]interface{}{"id": itemId.GetLocalID()})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		if err == appCommon.RecordNotFound {
			return nil, appCommon.ErrEntityNotFound(itemmodel.EntityName, err)
		}
		return nil, appCommon.ErrCannotGetEntity(itemmodel.EntityName, err)
	}

	cart, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"user_id": user.Id,
		"item_id": itemId.GetLocalID(),
	})

	if err != nil {
		if err == appCommon.RecordNotFound {
			if item.Quantity < data.Quantity {
				return nil, cartmodel.ErrQuantityExceed
			}
			createData := cartmodel.Cart{
				UserId:   user.Id,
				ItemId:   int64(itemId.GetLocalID()),
				Quantity: data.Quantity,
			}
			if err := biz.store.Create(ctx, &createData); err != nil {
				return nil, appCommon.ErrCannotCreateEntity(cartmodel.EntityName, err)
			}
			createData.Mask(false)
			return &createData, nil
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(cartmodel.EntityName, err)
	}

	if cart.Quantity+data.Quantity > item.Quantity {
		return nil, cartmodel.ErrQuantityExceed
	}

	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{
		"id": cart.Id,
	}, map[string]interface{}{
		"quantity": cart.Quantity + data.Quantity,
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotUpdateEntity(cartmodel.EntityName, err)
	}
	cart.Quantity += data.Quantity
	cart.Mask(false)
	return cart, nil
}
