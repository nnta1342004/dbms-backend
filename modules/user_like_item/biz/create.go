package userlikeitembiz

import (
	"context"
	"errors"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
	itemmodel "hareta/modules/item/model"
	usermodel "hareta/modules/user/model"
	userlikeitemmodel "hareta/modules/user_like_item/model"
)

type likeStore interface {
	Create(ctx context.Context, data *userlikeitemmodel.UserLikeItem) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*userlikeitemmodel.UserLikeItem, error)
}
type groupLikeStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*groupitemmodel.GroupItem, error)
}
type likeBiz struct {
	store      likeStore
	groupStore groupLikeStore
}

func NewLikeBiz(store likeStore, groupStore groupLikeStore) *likeBiz {
	return &likeBiz{store: store, groupStore: groupStore}
}

func (biz *likeBiz) Like(ctx context.Context, user *usermodel.User, data *userlikeitemmodel.UserLikeItemCreate) error {
	groupId, err := appCommon.FromBase58(data.GroupId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	res, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"user_id":  user.Id,
		"group_id": groupId.GetLocalID(),
	})

	if err != nil {
		if err != appCommon.RecordNotFound {
			return appCommon.ErrCannotGetEntity(userlikeitemmodel.EntityName, err)
		}
	}

	if res != nil {
		return appCommon.ErrEntityExisted(userlikeitemmodel.EntityName, errors.New("user has liked this item"))
	}

	_, err = biz.groupStore.FindDataWithCondition(ctx, map[string]interface{}{
		"id": groupId.GetLocalID(),
	})

	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(itemmodel.EntityName, err)
		}
		return appCommon.ErrCannotGetEntity(itemmodel.EntityName, err)
	}

	if err := biz.store.Create(ctx, &userlikeitemmodel.UserLikeItem{
		UserId:  user.Id,
		GroupId: int64(groupId.GetLocalID()),
	}); err != nil {
		return appCommon.ErrCannotCreateEntity(userlikeitemmodel.EntityName, err)
	}
	return nil
}
