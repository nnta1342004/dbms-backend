package passwordrecoverygin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	passwordrecoverybiz "hareta/modules/password_recovery/biz"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
	passwordrecoverystorage "hareta/modules/password_recovery/storage"
	userstorage "hareta/modules/user/storage"
	"hareta/plugin/pubsub"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data passwordrecoverymodel.UserRecreatePassword
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := passwordrecoverystorage.NewSQLStore(db)
		userStore := userstorage.NewSQLStore(db)
		rabbitMQ := sc.MustGet(appCommon.PluginRabbitMQ).(pubsub.PubSub)
		biz := passwordrecoverybiz.NewCreateBiz(store, userStore, rabbitMQ)
		_, err := biz.Create(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
