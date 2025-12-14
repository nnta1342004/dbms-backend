package uploadbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
	"time"
)

type listStore interface {
	List(
		ctx context.Context,
		paging *appCommon.Paging,
		filter *imagemodel.ImageFilter,
		conditions map[string]interface{},
		moreInfo ...string,
	) ([]imagemodel.Image, error)
}
type listBiz struct {
	store  listStore
	logger logger.Logger
}

func NewListBiz(store listStore) *listBiz {
	return &listBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("ImageListBiz"),
	}
}

func (biz *listBiz) ListDataWithCondition(ctx context.Context, filter *imagemodel.ImageList, paging *appCommon.Paging, moreInfo ...string) ([]imagemodel.Image, error) {
	filter.FulFill()
	paging.Fulfill()

	conditions := imagemodel.ImageFilter{
		TimeFrom: time.Unix(*filter.TimeFrom, 0),
		TimeTo:   time.Unix(*filter.TimeTo, 0),
	}

	result, err := biz.store.List(ctx, paging, &conditions, nil, moreInfo...)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []imagemodel.Image{}, appCommon.ErrCannotListEntity(imagemodel.EntityName, err)
	}

	for i := range result {
		result[i].Mask(false)
	}
	return result, nil
}
