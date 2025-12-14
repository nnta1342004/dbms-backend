package orderbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	ordermodel "hareta/modules/order/model"
	usermodel "hareta/modules/user/model"
)

type getOrderStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*ordermodel.Order, error)
}
type getOrderBiz struct {
	store  getOrderStore
	logger logger.Logger
}

func NewGetOrderBiz(store getOrderStore) *getOrderBiz {
	return &getOrderBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("GetOrderBiz"),
	}
}

func (biz *getOrderBiz) GetOrder(ctx context.Context, user *usermodel.User, filter *ordermodel.OrderFind) (*ordermodel.Order, error) {
	orderId, err := appCommon.FromBase58(filter.OrderId)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	mp := make(map[string]interface{})
	mp["id"] = orderId.GetLocalID()
	if user != nil {
		mp["user_id"] = user.Id
	}

	order, err := biz.store.FindDataWithCondition(ctx, mp)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(ordermodel.EntityName, err)
	}
	order.Mask(false)
	return order, nil
}
