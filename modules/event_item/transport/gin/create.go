package eventitemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	eventstorage "hareta/modules/event/storage"
	eventitembiz "hareta/modules/event_item/biz"
	eventitemmodel "hareta/modules/event_item/model"
	eventitemstorage "hareta/modules/event_item/storage"
	itemstorage "hareta/modules/item/storage"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data eventitemmodel.EventItemCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := eventitemstorage.NewSQLStore(db)
		eventStore := eventstorage.NewSQLStore(db)
		itemStore := itemstorage.NewSQLStore(db)
		biz := eventitembiz.NewCreateBiz(store, eventStore, itemStore)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
