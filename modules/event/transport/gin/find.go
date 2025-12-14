package eventgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	eventbiz "hareta/modules/event/biz"
	eventstorage "hareta/modules/event/storage"
	"net/http"
)

func Find(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := eventstorage.NewSQLStore(db)
		biz := eventbiz.NewFindEventBiz(store)
		res, err := biz.FindDataWithCondition(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
