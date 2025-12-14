package uploadbiz

import (
	"context"
	"fmt"
	"github.com/leductoan3082004/go-sdk/appCommon"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	imagemodel "hareta/modules/image/model"
	"path/filepath"
	"time"
)

type uploadStore interface {
	Create(ctx context.Context, data *imagemodel.Image) error
}
type uploadBiz struct {
	store uploadStore
	s3    aws.S3
}

func NewUploadBiz(store uploadStore, s3 aws.S3) *uploadBiz {
	return &uploadBiz{store: store, s3: s3}
}

func (biz *uploadBiz) UploadImage(ctx context.Context, data []byte, fileName string) (*imagemodel.Image, error) {
	fileExt := filepath.Ext(fileName) // "img.jpg" => ".jpg"
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		return nil, imagemodel.ErrInvalidImageFormat
	}
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg
	_, err := biz.s3.UploadFileData(ctx, data, "image/"+fileName)
	if err != nil {
		return nil, imagemodel.ErrCannotUploadImage
	}
	createData := imagemodel.Image{
		URL:      "https://asset.haretaworkshop.com/image/" + fileName,
		FileName: fileName,
	}
	if err := biz.store.Create(ctx, &createData); err != nil {
		return nil, appCommon.ErrCannotCreateEntity(imagemodel.EntityName, err)
	}
	createData.Mask(false)
	return &createData, nil
}
