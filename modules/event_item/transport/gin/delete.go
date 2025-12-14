package eventitemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	eventitembiz "hareta/modules/event_item/biz"
	eventitemmodel "hareta/modules/event_item/model"
	eventitemstorage "hareta/modules/event_item/storage"
	"net/http"
)

func Delete(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data eventitemmodel.EventItemDelete
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := eventitemstorage.NewSQLStore(db)
		biz := eventitembiz.NewDeleteBiz(store)
		if err := biz.Delete(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
