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

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data itemimagemodel.ItemList
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)

		store := itemimagestorage.NewSQLStore(db)
		biz := itemimagebiz.NewListBiz(store)
		res, err := biz.List(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, data.Paging, nil))
	}
}
