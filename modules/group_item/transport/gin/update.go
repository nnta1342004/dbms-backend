package groupitemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	groupitembiz "hareta/modules/group_item/biz"
	groupitemmodel "hareta/modules/group_item/model"
	groupitemstorage "hareta/modules/group_item/storage"
	"net/http"
)

func Update(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data groupitemmodel.GroupUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := groupitemstorage.NewSQLStore(db)
		biz := groupitembiz.NewUpdateBiz(store)
		if err := biz.Update(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
