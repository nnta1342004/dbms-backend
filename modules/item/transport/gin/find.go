package itemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	itembiz "hareta/modules/item/biz"
	itemmodel "hareta/modules/item/model"
	itemstorage "hareta/modules/item/storage"
	"net/http"
)

func Find(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := itemstorage.NewSQLStore(db)
		biz := itembiz.NewFindBiz(store)
		res, err := biz.Find(c.Request.Context(), &itemmodel.ItemFind{Id: id})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
