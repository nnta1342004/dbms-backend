package itemimagebiz

import (
	"context"
	"fmt"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
	itemmodel "hareta/modules/item/model"
	itemimagemodel "hareta/modules/item_image/model"
	"path/filepath"
	"time"
)

type addStore interface {
	Create(ctx context.Context, data *itemimagemodel.ItemImage) error
}
type addImageStore interface {
	Create(ctx context.Context, data *imagemodel.Image) error
}
type addItemStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (
		*itemmodel.Item, error,
	)
}
type addBiz struct {
	store     addStore
	imgStore  addImageStore
	s3        aws.S3
	itemStore addItemStore
}

func NewAddBiz(store addStore, imgStore addImageStore, s3 aws.S3, itemStore addItemStore) *addBiz {
	return &addBiz{store: store, s3: s3, imgStore: imgStore, itemStore: itemStore}
}
func (biz *addBiz) Add(
	ctx context.Context, data []byte, fileName string, query *itemimagemodel.ItemCreate,
) (*itemimagemodel.ItemImage, error) {

	id, err := appCommon.FromBase58(query.ItemId)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}
	_, err = biz.itemStore.FindDataWithCondition(
		ctx, map[string]interface{}{
			"id": id.GetLocalID(),
		},
	)
	if err != nil {
		if err == appCommon.RecordNotFound {
			return nil, appCommon.ErrEntityNotFound(itemmodel.EntityName, err)
		}
		return nil, appCommon.ErrCannotGetEntity(itemmodel.EntityName, err)
	}
	fileExt := filepath.Ext(fileName) // "img.jpg" => ".jpg"
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		return nil, imagemodel.ErrInvalidImageFormat
	}
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg
	_, err = biz.s3.UploadFileData(ctx, data, "image/"+fileName)

	if err != nil {
		return nil, appCommon.ErrInternal(err)
	}

	createData := imagemodel.Image{
		URL:      "https://asset.haretaworkshop.com/image/" + fileName,
		FileName: fileName,
	}

	if err := biz.imgStore.Create(ctx, &createData); err != nil {
		return nil, appCommon.ErrCannotCreateEntity(imagemodel.EntityName, err)
	}
	createData.Mask(false)
	dataCreate := itemimagemodel.ItemImage{
		ItemId:  int64(id.GetLocalID()),
		ImageId: createData.Id,
		Color:   query.Color,
		Image:   &createData,
	}
	if err := biz.store.Create(ctx, &dataCreate); err != nil {
		return nil, appCommon.ErrCannotCreateEntity(itemimagemodel.EntityName, err)
	}
	dataCreate.Mask(false)
	return &dataCreate, nil
}
