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
	"net/http"
)

func RecoverPassword(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data passwordrecoverymodel.RecoverData
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := passwordrecoverystorage.NewSQLStore(db)
		userStore := userstorage.NewSQLStore(db)
		biz := passwordrecoverybiz.NewRecoverBiz(store, userStore)
		db = db.Begin()
		if err := biz.Recover(c.Request.Context(), &data); err != nil {
			db.Rollback()
			panic(err)
		}
		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
