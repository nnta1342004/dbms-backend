package imagegin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	uploadbiz "hareta/modules/image/biz"
	imagemodel "hareta/modules/image/model"
	imagestorage "hareta/modules/image/storage"
	"net/http"
)

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging appCommon.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		var filter imagemodel.ImageList
		if err := c.ShouldBind(&filter); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := imagestorage.NewSQLStore(db)
		biz := uploadbiz.NewListBiz(store)
		res, err := biz.ListDataWithCondition(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
