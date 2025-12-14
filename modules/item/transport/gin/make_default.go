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

func MakeDefault(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data itemmodel.ItemMakeDefault
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		db = db.Begin()
		store := itemstorage.NewSQLStore(db)
		biz := itembiz.NewMakeDefaultBiz(store)
		if err := biz.MakeDefault(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}
		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
