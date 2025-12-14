package orderbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	ordermodel "hareta/modules/order/model"
)

type getOrderAdminStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*ordermodel.Order, error)
}
type getOrderAdminBiz struct {
	store  getOrderAdminStore
	logger logger.Logger
}

func NewGetOrderAdminBiz(store getOrderAdminStore) *getOrderAdminBiz {
	return &getOrderAdminBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("GetOrderAdminBiz"),
	}
}

func (biz *getOrderAdminBiz) GetOrder(ctx context.Context, filter *ordermodel.OrderFind) (*ordermodel.Order, error) {
	orderId, err := appCommon.FromBase58(filter.OrderId)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	order, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"id": orderId.GetLocalID(),
	})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(ordermodel.EntityName, err)
	}
	order.Mask(false)
	return order, nil
}
