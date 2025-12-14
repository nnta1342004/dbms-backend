package imagegin

import (
	"errors"
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	"gorm.io/gorm"
	"hareta/appCommon"
	uploadbiz "hareta/modules/image/biz"
	imagestorage "hareta/modules/image/storage"
	"net/http"
)

func UploadByFile(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		file, err := fileHeader.Open()
		if err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		defer file.Close()

		if fileHeader.Size > int64(1024*1024*15) {
			panic(appCommon.ErrInvalidRequest(errors.New("file size too large")))
		}
		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		s3 := sc.MustGet(appCommon.PluginAws).(aws.S3)

		imageStore := imagestorage.NewSQLStore(db)
		imageBiz := uploadbiz.NewUploadBiz(imageStore, s3)

		res, err := imageBiz.UploadImage(c.Request.Context(), dataBytes, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
