package groupitembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
	itemmodel "hareta/modules/item/model"
)

type deleteStore interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
}
type itemDeleteStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}
type deleteBiz struct {
	store     deleteStore
	logger    logger.Logger
	itemStore itemDeleteStore
}

func NewDeleteBiz(store deleteStore, itemStore itemDeleteStore) *deleteBiz {
	return &deleteBiz{
		store:     store,
		itemStore: itemStore,
		logger:    logger.GetCurrent().GetLogger("GroupItemDeleteBiz"),
	}
}

func (biz *deleteBiz) Delete(ctx context.Context, data *groupitemmodel.GroupDelete) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	if err := biz.store.DeleteWithCondition(ctx, map[string]interface{}{
		"id": id.GetLocalID(),
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(groupitemmodel.EntityName, err)
	}
	if err := biz.itemStore.UpdateWithCondition(ctx, map[string]interface{}{
		"group_id": id.GetLocalID(),
	}, map[string]interface{}{
		"status": itemmodel.StatusDeleted,
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(itemmodel.EntityName, err)
	}
	return nil
}
