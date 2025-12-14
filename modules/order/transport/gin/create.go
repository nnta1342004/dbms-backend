package ordergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	cartstorage "hareta/modules/cart/storage"
	itemstorage "hareta/modules/item/storage"
	itemorderstorage "hareta/modules/item_order/storage"
	orderbiz "hareta/modules/order/biz"
	ordermodel "hareta/modules/order/model"
	orderstorage "hareta/modules/order/storage"
	usermodel "hareta/modules/user/model"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ordermodel.OrderCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		db = db.Begin()

		store := orderstorage.NewSQLStore(db)
		cartStore := cartstorage.NewSQLStore(db)
		itemStore := itemstorage.NewSQLStore(db)
		itemOrderStore := itemorderstorage.NewSQLStore(db)

		biz := orderbiz.NewCreateBiz(store, cartStore, itemStore, itemOrderStore)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		res, err := biz.Create(c.Request.Context(), user, &data)
		if err != nil {
			db.Rollback()
			panic(err)
		}

		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
