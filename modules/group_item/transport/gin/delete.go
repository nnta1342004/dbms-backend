package groupitemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	groupitembiz "hareta/modules/group_item/biz"
	groupitemmodel "hareta/modules/group_item/model"
	groupitemstorage "hareta/modules/group_item/storage"
	itemstorage "hareta/modules/item/storage"
	"net/http"
)

func Delete(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data groupitemmodel.GroupDelete
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := groupitemstorage.NewSQLStore(db)
		itemStore := itemstorage.NewSQLStore(db)
		biz := groupitembiz.NewDeleteBiz(store, itemStore)
		db = db.Begin()
		if err := biz.Delete(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}
		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
