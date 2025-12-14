package orderbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	ordermodel "hareta/modules/order/model"
)

type listOrderAdminStore interface {
	List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]ordermodel.Order, error)
}
type listOrderAdminBiz struct {
	logger logger.Logger
	store  listOrderStore
}

func NewListOrderAdminBiz(store listOrderAdminStore) *listOrderAdminBiz {
	return &listOrderAdminBiz{logger: logger.GetCurrent().GetLogger("ListOrderAdminBiz"), store: store}
}
func (biz *listOrderAdminBiz) List(ctx context.Context, paging *appCommon.Paging, conditions *ordermodel.OrderListAdmin) ([]ordermodel.Order, error) {
	paging.Fulfill()
	condition := make(map[string]interface{})
	if conditions.Status != nil {
		condition["status"] = *conditions.Status
	}
	if conditions.Email != nil {
		condition["email"] = *conditions.Email
	}
	orders, err := biz.store.List(ctx, paging, condition)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	for i := range orders {
		orders[i].Mask(false)
	}
	return orders, nil
}
