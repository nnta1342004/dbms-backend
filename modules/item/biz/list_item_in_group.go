package itembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

type listGroupStore interface {
	ListItemInGroup(
		ctx context.Context,
		paging *appCommon.Paging,
		groupId int64,
		moreInfo ...string,
	) ([]itemmodel.SimpleItem, error)
}

type listGroupBiz struct {
	store  listGroupStore
	logger logger.Logger
}

func NewListGroupBiz(store listGroupStore) *listGroupBiz {
	return &listGroupBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ItemListGroupBiz"),
	}
}

func (biz *listGroupBiz) ListItemInGroup(
	ctx context.Context, paging *appCommon.Paging, filter *itemmodel.ItemListGroup,
) ([]itemmodel.SimpleItem, error) {
	id, err := appCommon.FromBase58(filter.Id)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}

	res, err := biz.store.ListItemInGroup(ctx, paging, int64(id.GetLocalID()), "Group", "Avatar")
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(itemmodel.EntityName, err)
	}
	for i := range res {
		res[i].Mask(false)
		if res[i].Avatar != nil {
			res[i].Avatar.Mask(false)
		}
		res[i].Group.Mask(false)
	}
	if res != nil && len(res) > 0 {
		paging.NextCursor = res[len(res)-1].FakeId.String()
	}

	return res, nil
}
