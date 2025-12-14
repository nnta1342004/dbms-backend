package usergin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	userbiz "hareta/modules/user/biz"
	usermodel "hareta/modules/user/model"
	userstorage "hareta/modules/user/storage"
	"net/http"
)

func ChangePassword(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserChangePassword
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		rdc := sc.MustGet(appCommon.PluginRedis).(*redis.Client)
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewChangePasswordBiz(store)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)

		if err := biz.ChangePassword(c.Request.Context(), user, &data, rdc); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
