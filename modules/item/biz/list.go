package itembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

type listStore interface {
	List(
		ctx context.Context,
		filter *itemmodel.ItemList,
		paging *appCommon.Paging,
		moreInfo ...string,
	) ([]itemmodel.SimpleItem, error)
}
type listDefaultItemBiz struct {
	store  listStore
	logger logger.Logger
}

func NewListDefaultItemBiz(store listStore) *listDefaultItemBiz {
	return &listDefaultItemBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ListDefaultItemBiz"),
	}
}

func (biz *listDefaultItemBiz) ListDefaultItem(
	ctx context.Context,
	paging *appCommon.Paging,
	filter *itemmodel.ItemList,
) ([]itemmodel.SimpleItem, error) {
	filter.FulFill()
	paging.Fulfill()

	res, err := biz.store.List(ctx, filter, paging, "Group", "Avatar")
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []itemmodel.SimpleItem{}, appCommon.ErrCannotListEntity(itemmodel.EntityName, err)
	}

	for i, item := range res {
		res[i].Mask(false)
		if item.Avatar != nil {
			item.Avatar.Mask(false)
		}
		if item.Group != nil {
			item.Group.Mask(false)
		}
	}

	return res, nil
}
