package usergin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/tokenprovider"
	"gorm.io/gorm"
	"hareta/appCommon"
	userbiz "hareta/modules/user/biz"
	usermodel "hareta/modules/user/model"
	userstorage "hareta/modules/user/storage"
	"net/http"
)

func Login(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		rdc := sc.MustGet(appCommon.PluginRedis).(*redis.Client)
		tokenProvider := sc.MustGet(appCommon.PluginJwt).(tokenprovider.Provider)
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewLoginBiz(store)
		token, err := biz.Login(c.Request.Context(), &data, tokenProvider, rdc)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(token))
	}
}
