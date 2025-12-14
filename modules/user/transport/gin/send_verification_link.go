package usergin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	goservice "github.com/leductoan3082004/go-sdk"
	mailplugin "github.com/leductoan3082004/go-sdk/plugin/mailer"
	"gorm.io/gorm"
	"hareta/appCommon"
	userbiz "hareta/modules/user/biz"
	usermodel "hareta/modules/user/model"
	userstorage "hareta/modules/user/storage"
	"net/http"
)

func SendLink(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLink
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		mail := sc.MustGet(appCommon.PluginMailer).(mailplugin.MailEngine)
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		rdc := sc.MustGet(appCommon.PluginRedis).(*redis.Client)
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewSendLinkBiz(store)
		if err := biz.SendLink(c.Request.Context(), data.Email, mail, rdc); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
