package blogtaggin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	blogtaginfra "hareta/modules/blog-tag/infra"
	"net/http"
)

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := blogtaginfra.NewSQLStore(db)
		res, err := store.ListTags(c.Request.Context())
		if err != nil {
			panic(appCommon.ErrInternal(err))
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
