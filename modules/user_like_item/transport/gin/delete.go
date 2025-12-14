package userlikeitemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
	userlikeitembiz "hareta/modules/user_like_item/biz"
	userlikeitemmodel "hareta/modules/user_like_item/model"
	userlikeitemstorage "hareta/modules/user_like_item/storage"
	"net/http"
)

func Delete(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userlikeitemmodel.UserLikeItemDelete
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)

		store := userlikeitemstorage.NewSQLStore(db)
		biz := userlikeitembiz.NewDeleteBiz(store)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)

		if err := biz.Delete(c.Request.Context(), user, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
