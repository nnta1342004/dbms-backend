package uploadbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
)

type deleteStore interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*imagemodel.Image, error)
}

type deleteBiz struct {
	store  deleteStore
	s3     aws.S3
	logger logger.Logger
}

func NewDeleteBiz(store deleteStore, s3 aws.S3) *deleteBiz {
	return &deleteBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("DeleteImageBiz"),
		s3:     s3,
	}
}

func (biz *deleteBiz) DeleteImage(ctx context.Context, filter *imagemodel.ImageDelete) error {
	id, err := appCommon.FromBase58(filter.ImageId)
	if err != nil {
		return appCommon.ErrInvalidRequest(err)
	}

	image, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"id": id.GetLocalID(),
	})

	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(imagemodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(imagemodel.EntityName, err)
	}

	path, err := appCommon.GetPathFromUrl(image.URL)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}

	if err := biz.s3.DeleteImages(ctx, []string{path}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}

	if err := biz.store.DeleteWithCondition(ctx, map[string]interface{}{
		"id": id.GetLocalID(),
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(imagemodel.EntityName, err)
	}
	return nil
}
