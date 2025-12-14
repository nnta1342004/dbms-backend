package itemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"gorm.io/gorm"
	"hareta/appCommon"
	uploadbiz "hareta/modules/image/biz"
	imagestorage "hareta/modules/image/storage"
	itembiz "hareta/modules/item/biz"
	itemmodel "hareta/modules/item/model"
	itemstorage "hareta/modules/item/storage"
	"net/http"
)

func UpdateAvt(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		id := c.PostForm("id")
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		img, err := file.Open()
		defer img.Close()
		dataBytes := make([]byte, file.Size)
		if _, err := img.Read(dataBytes); err != nil {
			panic(appCommon.ErrInternal(err))
		}

		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		s3 := sc.MustGet(appCommon.PluginAws).(aws.S3)
		imgStore := imagestorage.NewSQLStore(db)
		itemStore := itemstorage.NewSQLStore(db)

		db = db.Begin()
		imgBiz := uploadbiz.NewUploadBiz(imgStore, s3)
		image, err := imgBiz.UploadImage(c.Request.Context(), dataBytes, file.Filename)
		if err != nil {
			panic(err)
		}
		itemBiz := itembiz.NewUpdateAvtBiz(itemStore, imgStore, s3)
		if err := itemBiz.UpdateAvt(c.Request.Context(), &itemmodel.ItemAvtUpdate{
			Id:      id,
			ImageId: image.Id,
		}); err != nil {
			db.Rollback()
			panic(err)
		}
		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
