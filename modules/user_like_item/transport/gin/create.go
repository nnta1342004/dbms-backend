package userlikeitemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	groupitemstorage "hareta/modules/group_item/storage"
	usermodel "hareta/modules/user/model"
	userlikeitembiz "hareta/modules/user_like_item/biz"
	userlikeitemmodel "hareta/modules/user_like_item/model"
	userlikeitemstorage "hareta/modules/user_like_item/storage"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userlikeitemmodel.UserLikeItemCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		db = db.Begin()
		store := userlikeitemstorage.NewSQLStore(db)
		groupStore := groupitemstorage.NewSQLStore(db)
		biz := userlikeitembiz.NewLikeBiz(store, groupStore)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)

		if err := biz.Like(c.Request.Context(), user, &data); err != nil {
			db.Rollback()
			panic(err)
		}
		if err := db.Commit().Error; err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
