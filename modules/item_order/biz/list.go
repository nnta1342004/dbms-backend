package itemorderbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	itemordermodel "hareta/modules/item_order/model"
	ordermodel "hareta/modules/order/model"
	usermodel "hareta/modules/user/model"
)

type itemOrderStore interface {
	List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]itemordermodel.ItemOrder, error)
}
type orderStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*ordermodel.Order, error)
}
type itemOrderBiz struct {
	store      itemOrderStore
	logger     logger.Logger
	orderStore orderStore
}

func NewItemOrderBiz(store itemOrderStore, orderStore orderStore) *itemOrderBiz {
	return &itemOrderBiz{
		store:      store,
		logger:     logger.GetCurrent().GetLogger("ListItemOrderBiz"),
		orderStore: orderStore,
	}
}

func (biz *itemOrderBiz) List(ctx context.Context, user *usermodel.User, paging *appCommon.Paging, filter *itemordermodel.ItemOrderList) ([]itemordermodel.ItemOrder, error) {
	paging.Fulfill()
	orderId, err := appCommon.FromBase58(filter.OrderId)
	if err != nil {
		return []itemordermodel.ItemOrder{}, appCommon.ErrInvalidRequest(err)
	}

	userId := int64(0)
	if user != nil {
		userId = user.Id
	}
	_, err = biz.orderStore.FindDataWithCondition(ctx, map[string]interface{}{
		"id":      orderId.GetLocalID(),
		"user_id": userId,
	})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		if err == appCommon.RecordNotFound {
			return []itemordermodel.ItemOrder{}, appCommon.ErrEntityNotFound(ordermodel.EntityName, err)
		}
		return []itemordermodel.ItemOrder{}, appCommon.ErrCannotGetEntity(ordermodel.EntityName, err)
	}

	res, err := biz.store.List(ctx, paging, map[string]interface{}{
		"order_id": orderId.GetLocalID(),
	}, "Item", "Item.Avatar")
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []itemordermodel.ItemOrder{}, appCommon.ErrCannotListEntity(itemordermodel.EntityName, err)
	}
	if res == nil {
		return []itemordermodel.ItemOrder{}, nil
	}
	for i := range res {
		res[i].Mask(false)
		if res[i].Item != nil {
			res[i].Item.Mask(false)
			if res[i].Item.Avatar != nil {
				res[i].Item.Avatar.Mask(false)
			}
			if res[i].Item.Group != nil {
				res[i].Item.Group.Mask(false)
			}
		}
	}

	if len(res) > 0 {
		paging.NextCursor = res[len(res)-1].FakeId.String()
	}
	return res, nil
}
