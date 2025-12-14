package orderbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	ordermodel "hareta/modules/order/model"
	usermodel "hareta/modules/user/model"
)

type listOrderStore interface {
	List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]ordermodel.Order, error)
}
type listOrderBiz struct {
	store  listOrderStore
	logger logger.Logger
}

func NewListOrderBiz(store listOrderStore) *listOrderBiz {
	return &listOrderBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ListOrderBiz"),
	}
}

func (biz *listOrderBiz) ListOrder(ctx context.Context, paging *appCommon.Paging, user *usermodel.User) ([]ordermodel.Order, error) {
	paging.Fulfill()
	orders, err := biz.store.List(ctx, paging, map[string]interface{}{"user_id": user.Id})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	for i := range orders {
		orders[i].Mask(false)
	}
	return orders, nil
}
