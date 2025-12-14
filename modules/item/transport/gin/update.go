package itemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	imagestorage "hareta/modules/image/storage"
	itembiz "hareta/modules/item/biz"
	itemmodel "hareta/modules/item/model"
	itemstorage "hareta/modules/item/storage"
	"net/http"
)

func Update(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data itemmodel.ItemUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := itemstorage.NewSQLStore(db)
		imgStore := imagestorage.NewSQLStore(db)
		biz := itembiz.NewUpdateBiz(store, imgStore)
		if err := biz.Update(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
