package itemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	groupitemstorage "hareta/modules/group_item/storage"
	itembiz "hareta/modules/item/biz"
	itemmodel "hareta/modules/item/model"
	itemstorage "hareta/modules/item/storage"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data itemmodel.ItemCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := itemstorage.NewSQLStore(db)
		groupStore := groupitemstorage.NewSQLStore(db)
		biz := itembiz.NewCreateBiz(store, groupStore)
		res, err := biz.Create(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
