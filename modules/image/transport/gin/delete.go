package imagegin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"gorm.io/gorm"
	"hareta/appCommon"
	uploadbiz "hareta/modules/image/biz"
	imagemodel "hareta/modules/image/model"
	imagestorage "hareta/modules/image/storage"
	"net/http"
)

func Delete(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data imagemodel.ImageDelete
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		db = db.Begin()
		store := imagestorage.NewSQLStore(db)
		s3 := sc.MustGet(appCommon.PluginAws).(aws.S3)
		biz := uploadbiz.NewDeleteBiz(store, s3)
		if err := biz.DeleteImage(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}
		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
