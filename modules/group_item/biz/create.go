package groupitembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
)

type createStore interface {
	Create(ctx context.Context, data *groupitemmodel.GroupItem) error
}
type createBiz struct {
	logger logger.Logger
	store  createStore
}

func NewCreateBiz(store createStore) *createBiz {
	return &createBiz{logger: logger.GetCurrent().GetLogger("GroupCreateBiz"), store: store}
}

func (biz *createBiz) Create(ctx context.Context, data *groupitemmodel.GroupCreate) (*groupitemmodel.GroupItem, error) {
	createData := &groupitemmodel.GroupItem{
		Name: data.Name,
	}
	if err := biz.store.Create(ctx, createData); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(groupitemmodel.EntityName, err)
	}
	createData.Mask(false)
	return createData, nil
}
