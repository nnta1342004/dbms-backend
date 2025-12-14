package cartbiz

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
	usermodel "hareta/modules/user/model"
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

func (biz *deleteBiz) Delete(ctx context.Context, user *usermodel.User, data *cartmodel.CartDelete) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	if err := biz.store.DeleteWithCondition(ctx, map[string]interface{}{
		"id":      id.GetLocalID(),
		"user_id": user.Id,
	}); err != nil {
		return appCommon.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}
	return nil
}
