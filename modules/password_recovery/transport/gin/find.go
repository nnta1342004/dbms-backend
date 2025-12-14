package passwordrecoverygin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	passwordrecoverybiz "hareta/modules/password_recovery/biz"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
	passwordrecoverystorage "hareta/modules/password_recovery/storage"
	"net/http"
)

func Find(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data passwordrecoverymodel.PasswordRecoveryFind
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := passwordrecoverystorage.NewSQLStore(db)
		biz := passwordrecoverybiz.NewFindBiz(store)
		res, err := biz.Find(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
