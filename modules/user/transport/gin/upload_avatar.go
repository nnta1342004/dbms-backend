package usergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"gorm.io/gorm"
	"hareta/appCommon"
	imagestorage "hareta/modules/image/storage"
	userbiz "hareta/modules/user/biz"
	usermodel "hareta/modules/user/model"
	userstorage "hareta/modules/user/storage"
	"net/http"
)

func UploadAvatar(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
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
		s3 := sc.MustGet(appCommon.PluginAws).(aws.S3)
		db = db.Begin()
		store := userstorage.NewSQLStore(db)
		imgStore := imagestorage.NewSQLStore(db)
		biz := userbiz.NewUploadBiz(store, imgStore, s3)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		if err := biz.Upload(c.Request.Context(), user, dataBytes, file.Filename); err != nil {
			db.Rollback()
			panic(err)
		}
		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
