package userlikeitembiz

import (
	"context"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
	userlikeitemmodel "hareta/modules/user_like_item/model"
)

type deleteStore interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
}
type deleteBiz struct {
	store deleteStore
}

func NewDeleteBiz(store deleteStore) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) Delete(ctx context.Context, user *usermodel.User, data *userlikeitemmodel.UserLikeItemDelete) error {
	id, err := appCommon.FromBase58(data.GroupId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	if err := biz.store.DeleteWithCondition(ctx, map[string]interface{}{
		"user_id":  user.Id,
		"group_id": id.GetLocalID(),
	}); err != nil {
		return appCommon.ErrCannotDeleteEntity(userlikeitemmodel.EntityName, err)
	}
	return nil
}
