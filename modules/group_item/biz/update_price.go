package groupitembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
	itemmodel "hareta/modules/item/model"
)

type updatePriceStore interface {
	UpdateWithCondition(
		ctx context.Context,
		conditions map[string]interface{},
		data map[string]interface{},
	) error
}

type updatePriceBiz struct {
	store  updatePriceStore
	logger logger.Logger
}

func NewUpdatePriceBiz(store updatePriceStore) *updatePriceBiz {
	return &updatePriceBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("UpdatePriceBiz"),
	}
}

func (biz *updatePriceBiz) UpdatePrice(ctx context.Context, priceID string, data *groupitemmodel.GroupPriceUpdate) error {
	priceId, err := appCommon.FromBase58(priceID)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	mp := make(map[string]interface{})

	if data.OriginalPrice != nil {
		mp["original_price"] = *data.OriginalPrice
	}
	if data.Price != nil {
		mp["price"] = *data.Price
	}

	if err := biz.store.UpdateWithCondition(
		ctx,
		map[string]interface{}{"group_id": priceId.GetLocalID()},
		mp,
	); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(itemmodel.EntityName, err)
	}
	return nil
}
