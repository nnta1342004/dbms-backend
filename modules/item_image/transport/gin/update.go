package itemimagegin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	itemimagebiz "hareta/modules/item_image/biz"
	itemimagemodel "hareta/modules/item_image/model"
	itemimagestorage "hareta/modules/item_image/storage"
	"net/http"
)

func Update(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data itemimagemodel.ItemUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)

		store := itemimagestorage.NewSQLStore(db)
		biz := itemimagebiz.NewUpdateBiz(store)
		if err := biz.Update(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
