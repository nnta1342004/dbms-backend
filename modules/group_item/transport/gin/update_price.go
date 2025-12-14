package groupitemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	groupitembiz "hareta/modules/group_item/biz"
	groupitemmodel "hareta/modules/group_item/model"
	itemstorage "hareta/modules/item/storage"
	"net/http"
)

func UpdatePrice(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data groupitemmodel.GroupPriceUpdate
		groupID := c.Param("group-id")
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := itemstorage.NewSQLStore(db)
		biz := groupitembiz.NewUpdatePriceBiz(store)
		if err := biz.UpdatePrice(c.Request.Context(), groupID, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
