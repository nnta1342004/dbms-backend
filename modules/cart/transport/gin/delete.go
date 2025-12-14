package cartgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	cartbiz "hareta/modules/cart/biz"
	cartmodel "hareta/modules/cart/model"
	cartstorage "hareta/modules/cart/storage"
	usermodel "hareta/modules/user/model"
	"net/http"
)

func Delete(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data cartmodel.CartDelete
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := cartstorage.NewSQLStore(db)
		biz := cartbiz.NewDeleteBiz(store)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		err := biz.Delete(c.Request.Context(), user, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
