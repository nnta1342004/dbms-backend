package itemimagegin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"gorm.io/gorm"
	"hareta/appCommon"
	imagestorage "hareta/modules/image/storage"
	itemstorage "hareta/modules/item/storage"
	itemimagebiz "hareta/modules/item_image/biz"
	itemimagemodel "hareta/modules/item_image/model"
	itemimagestorage "hareta/modules/item_image/storage"
	"net/http"
)

func AddImages(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		var data itemimagemodel.ItemCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		img, err := file.Open()
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		defer img.Close()
		dataBytes := make([]byte, file.Size)
		if _, err := img.Read(dataBytes); err != nil {
			panic(appCommon.ErrInternal(err))
		}

		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		db = db.Begin()
		store := itemimagestorage.NewSQLStore(db)
		imgStore := imagestorage.NewSQLStore(db)
		s3 := sc.MustGet(appCommon.PluginAws).(aws.S3)
		itemStore := itemstorage.NewSQLStore(db)
		biz := itemimagebiz.NewAddBiz(store, imgStore, s3, itemStore)
		res, err := biz.Add(c.Request.Context(), dataBytes, file.Filename, &data)

		if err != nil {
			db.Rollback()
			panic(err)
		}
		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
