package cartgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	cartbiz "hareta/modules/cart/biz"
	cartstorage "hareta/modules/cart/storage"
	usermodel "hareta/modules/user/model"
	"net/http"
)

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging appCommon.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := cartstorage.NewSQLStore(db)
		biz := cartbiz.NewListBiz(store)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		res, err := biz.List(c.Request.Context(), user, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
