package ordergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	orderbiz "hareta/modules/order/biz"
	ordermodel "hareta/modules/order/model"
	orderstorage "hareta/modules/order/storage"
	usermodel "hareta/modules/user/model"
	"net/http"
)

func Find(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("id")
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewGetOrderBiz(store)

		res, err := biz.GetOrder(c.Request.Context(), user, &ordermodel.OrderFind{OrderId: orderId})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
