package itemimagebiz

import (
	"context"
	"hareta/appCommon"
	itemimagemodel "hareta/modules/item_image/model"
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

func (biz *deleteBiz) Delete(ctx context.Context, data *itemimagemodel.ItemDelete) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	if err := biz.store.DeleteWithCondition(ctx, map[string]interface{}{"id": id.GetLocalID()}); err != nil {
		return appCommon.ErrCannotDeleteEntity(itemimagemodel.EntityName, err)
	}
	return nil
}
