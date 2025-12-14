package itemordergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	itemorderbiz "hareta/modules/item_order/biz"
	itemordermodel "hareta/modules/item_order/model"
	itemorderstorage "hareta/modules/item_order/storage"
	orderstorage "hareta/modules/order/storage"
	usermodel "hareta/modules/user/model"
	"net/http"
)

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data itemordermodel.ItemOrderList
		var paging appCommon.Paging
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := itemorderstorage.NewSQLStore(db)
		orderStore := orderstorage.NewSQLStore(db)
		biz := itemorderbiz.NewItemOrderBiz(store, orderStore)
		res, err := biz.List(c.Request.Context(), user, &paging, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
