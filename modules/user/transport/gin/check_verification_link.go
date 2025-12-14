package usergin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	userbiz "hareta/modules/user/biz"
	userstorage "hareta/modules/user/storage"
	"net/http"
)

func CheckLink(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		link := c.Param("id")
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)

		rdc := sc.MustGet(appCommon.PluginRedis).(*redis.Client)
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewCheckLinkBiz(store)
		if err := biz.CheckLink(c.Request.Context(), rdc, link); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
