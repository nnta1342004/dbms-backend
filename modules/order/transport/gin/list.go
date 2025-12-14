package ordergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	orderbiz "hareta/modules/order/biz"
	orderstorage "hareta/modules/order/storage"
	usermodel "hareta/modules/user/model"
	"net/http"
)

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging appCommon.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := orderstorage.NewSQLStore(db)
		biz := orderbiz.NewListOrderBiz(store)
		res, err := biz.ListOrder(c.Request.Context(), &paging, user)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
