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

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data groupitemmodel.GroupCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := groupitemstorage.NewSQLStore(db)
		biz := groupitembiz.NewCreateBiz(store)
		res, err := biz.Create(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
