package userbiz

import (
	"context"
	"fmt"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
	usermodel "hareta/modules/user/model"
	"path/filepath"
	"time"
)

type uploadStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}
type uploadImageStore interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
	Create(ctx context.Context, data *imagemodel.Image) error
}
type uploadBiz struct {
	store    uploadStore
	imgStore uploadImageStore
	s3       aws.S3
}

func NewUploadBiz(store uploadStore, imgStore uploadImageStore, s3 aws.S3) *uploadBiz {
	return &uploadBiz{store: store, imgStore: imgStore, s3: s3}
}

func (biz *uploadBiz) Upload(ctx context.Context, user *usermodel.User, dataBytes []byte, fileName string) error {
	if user.Avatar != nil {
		if err := biz.s3.DeleteImages(ctx, []string{user.Avatar.FileName}); err != nil {
			return appCommon.ErrInternal(err)
		}
		if err := biz.imgStore.DeleteWithCondition(
			ctx, map[string]interface{}{
				"id": user.Avatar.Id,
			},
		); err != nil {
			return appCommon.ErrCannotDeleteEntity(imagemodel.EntityName, err)
		}
	}
	fileExt := filepath.Ext(fileName) // "img.jpg" => ".jpg"
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		return imagemodel.ErrInvalidImageFormat
	}
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg
	_, err := biz.s3.UploadFileData(ctx, dataBytes, "image/"+fileName)
	if err != nil {
		return appCommon.ErrInternal(err)
	}
	createData := imagemodel.Image{
		URL:      "https://asset.haretaworkshop.com/image/" + fileName,
		FileName: fileName,
	}
	if err := biz.imgStore.Create(ctx, &createData); err != nil {
		return appCommon.ErrCannotCreateEntity(imagemodel.EntityName, err)
	}
	if err := biz.store.UpdateWithCondition(
		ctx, map[string]interface{}{
			"id": user.Id,
		}, map[string]interface{}{
			"avatar_id": createData.Id,
		},
	); err != nil {
		return appCommon.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
