package groupitembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
)

type updateStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}

type updateBiz struct {
	store  updateStore
	logger logger.Logger
}

func NewUpdateBiz(store updateStore) *updateBiz {
	return &updateBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("GroupItemUpdateBiz"),
	}
}

func (biz *updateBiz) Update(ctx context.Context, data *groupitemmodel.GroupUpdate) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{
		"id": id.GetLocalID(),
	}, map[string]interface{}{
		"name": data.Name,
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(groupitemmodel.EntityName, err)
	}
	return nil
}
