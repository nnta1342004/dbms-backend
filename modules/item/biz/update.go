package itembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
	itemmodel "hareta/modules/item/model"
)

type updateStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}
type imgUpdateStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*imagemodel.Image, error)
}
type updateBiz struct {
	store    updateStore
	imgStore imgUpdateStore
	logger   logger.Logger
}

func NewUpdateBiz(store updateStore, imgStore imgStore) *updateBiz {
	return &updateBiz{store: store, imgStore: imgStore, logger: logger.GetCurrent().GetLogger("ItemUpdateBiz")}
}

func (biz *updateBiz) Update(ctx context.Context, data *itemmodel.ItemUpdate) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	updateData, err := data.ToMap()

	if err != nil {
		return err
	}

	if data.AvatarFakeId != nil {
		avtId, err := appCommon.FromBase58(*data.AvatarFakeId)

		if err != nil {
			return appCommon.ErrInvalidRequest(err)
		}
		_, err = biz.imgStore.FindDataWithCondition(ctx, map[string]interface{}{
			"id": avtId.GetLocalID(),
		})
		if err != nil {
			biz.logger.WithSrc().Errorln(err)
			if err == appCommon.RecordNotFound {
				return appCommon.ErrEntityNotFound(imagemodel.EntityName, err)
			}
			return appCommon.ErrCannotGetEntity(imagemodel.EntityName, err)
		}
		updateData["avatar_id"] = avtId.GetLocalID()
	}

	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{"id": id.GetLocalID()}, updateData); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(itemmodel.EntityName, err)
	}
	return nil
}
