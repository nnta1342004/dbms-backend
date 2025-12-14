package groupitembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
)

type listStore interface {
	List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]groupitemmodel.GroupItem, error)
}

type listBiz struct {
	store  listStore
	logger logger.Logger
}

func NewListBiz(store listStore) *listBiz {
	return &listBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("GroupListBiz"),
	}
}

func (biz *listBiz) List(ctx context.Context, paging *appCommon.Paging) ([]groupitemmodel.GroupItem, error) {
	paging.Fulfill()
	res, err := biz.store.List(ctx, paging, nil)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(groupitemmodel.EntityName, err)
	}
	for i := range res {
		res[i].Mask(false)
	}
	return res, nil
}
