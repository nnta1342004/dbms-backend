package cartgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	cartbiz "hareta/modules/cart/biz"
	cartmodel "hareta/modules/cart/model"
	cartstorage "hareta/modules/cart/storage"
	itemstorage "hareta/modules/item/storage"
	usermodel "hareta/modules/user/model"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data cartmodel.CartCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := cartstorage.NewSQLStore(db)
		itemStore := itemstorage.NewSQLStore(db)
		biz := cartbiz.NewCreateBiz(store, itemStore)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		res, err := biz.Create(c.Request.Context(), user, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
