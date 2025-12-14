package itembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
	itemmodel "hareta/modules/item/model"
)

type deleteStore interface {
	SoftDeleteById(ctx context.Context, id int64) error
}
type cartDeleteStore interface {
	SoftDeleteByItemId(ctx context.Context, itemId int64) error
}
type deleteBiz struct {
	logger    logger.Logger
	store     deleteStore
	cartStore cartDeleteStore
}

func NewDeleteBiz(store deleteStore, cartStore cartDeleteStore) *deleteBiz {
	return &deleteBiz{
		store:     store,
		logger:    logger.GetCurrent().GetLogger("CartDeleteBiz"),
		cartStore: cartStore,
	}
}

func (biz *deleteBiz) Delete(ctx context.Context, itemDelete *itemmodel.ItemDelete) error {
	id, err := appCommon.FromBase58(itemDelete.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	if err := biz.store.SoftDeleteById(ctx, int64(id.GetLocalID())); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(itemmodel.EntityName, err)
	}

	if err := biz.cartStore.SoftDeleteByItemId(ctx, int64(id.GetLocalID())); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}
	return nil
}
