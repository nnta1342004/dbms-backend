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

func ListItemInGroup(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data itemmodel.ItemListGroup
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		var paging appCommon.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := itemstorage.NewSQLStore(db)
		biz := itembiz.NewListGroupBiz(store)
		res, err := biz.ListItemInGroup(c.Request.Context(), &paging, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
