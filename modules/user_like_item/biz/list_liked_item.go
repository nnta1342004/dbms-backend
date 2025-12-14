package userlikeitembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
	usermodel "hareta/modules/user/model"
	userlikeitemmodel "hareta/modules/user_like_item/model"
)

type listLikedItemStore interface {
	List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]userlikeitemmodel.UserLikeItem, error)
}
type itemStore interface {
	ListDefaultItemInGroups(ctx context.Context, groupIds []int64, conditions map[string]interface{}, moreInfo ...string) ([]itemmodel.SimpleItem, error)
}
type listLikedItemBiz struct {
	store     listLikedItemStore
	itemStore itemStore
	logger    logger.Logger
}

func NewListLikedItemBiz(store listLikedItemStore, itemStore itemStore) *listLikedItemBiz {
	return &listLikedItemBiz{
		store:     store,
		logger:    logger.GetCurrent().GetLogger("ListLikedItemBiz"),
		itemStore: itemStore,
	}
}

func (biz *listLikedItemBiz) ListLikedItem(ctx context.Context, user *usermodel.User, paging *appCommon.Paging) ([]itemmodel.SimpleItem, error) {
	if paging == nil {
		paging = &appCommon.Paging{
			Page:  1,
			Limit: 50,
		}
	}
	res, err := biz.store.List(ctx, paging, map[string]interface{}{
		"user_id": user.Id,
	})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotListEntity(userlikeitemmodel.EntityName, err)
	}

	id := make([]int64, len(res))
	for i := range res {
		id[i] = res[i].GroupId
	}
	ans, err := biz.itemStore.ListDefaultItemInGroups(ctx, id, map[string]interface{}{
		"default": 1,
	}, "Avatar", "Group")
	if err != nil {
		return nil, appCommon.ErrCannotListEntity(itemmodel.EntityName, err)
	}
	for i := range ans {
		ans[i].Mask(false)
		if ans[i].Avatar != nil {
			ans[i].Avatar.Mask(false)
		}
		ans[i].Group.Mask(false)
	}
	return ans, nil
}
