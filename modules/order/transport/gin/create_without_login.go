package ordergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	itemstorage "hareta/modules/item/storage"
	itemorderstorage "hareta/modules/item_order/storage"
	orderbiz "hareta/modules/order/biz"
	ordermodel "hareta/modules/order/model"
	orderstorage "hareta/modules/order/storage"
	"net/http"
)

func CreateWithoutLogin(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ordermodel.OrderCreateWithoutLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := orderstorage.NewSQLStore(db)
		itemStore := itemstorage.NewSQLStore(db)
		itemOrderStore := itemorderstorage.NewSQLStore(db)
		db = db.Begin()
		biz := orderbiz.NewCreateWithoutLoginBiz(store, itemStore, itemOrderStore)
		res, err := biz.Create(c.Request.Context(), &data)
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
