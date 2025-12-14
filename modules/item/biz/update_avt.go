package itembiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
	itemmodel "hareta/modules/item/model"
)

type updateAvtStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*itemmodel.Item, error)
}
type imgStore interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*imagemodel.Image, error)
}
type updateAvtBiz struct {
	store    updateAvtStore
	imgStore imgStore
	s3       aws.S3
}

func NewUpdateAvtBiz(store updateAvtStore, imgStore imgStore, s3 aws.S3) *updateAvtBiz {
	return &updateAvtBiz{store: store, imgStore: imgStore, s3: s3}
}

func (biz *updateAvtBiz) UpdateAvt(ctx context.Context, data *itemmodel.ItemAvtUpdate) error {
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}
	_, err = biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id.GetLocalID()})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(itemmodel.EntityName, err)
		}
		return appCommon.ErrCannotGetEntity(itemmodel.EntityName, err)
	}

	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{
		"id": id.GetLocalID(),
	}, map[string]interface{}{
		"avatar_id": data.ImageId,
	}); err != nil {
		return appCommon.ErrCannotUpdateEntity(itemmodel.EntityName, err)
	}
	return nil
}
