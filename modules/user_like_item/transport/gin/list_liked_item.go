package userlikeitemgin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	itemstorage "hareta/modules/item/storage"
	usermodel "hareta/modules/user/model"
	userlikeitembiz "hareta/modules/user_like_item/biz"
	userlikeitemstorage "hareta/modules/user_like_item/storage"
	"net/http"
)

func ListLikedItem(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data appCommon.Paging
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := userlikeitemstorage.NewSQLStore(db)
		itemStore := itemstorage.NewSQLStore(db)
		biz := userlikeitembiz.NewListLikedItemBiz(store, itemStore)
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)

		res, err := biz.ListLikedItem(c.Request.Context(), user, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, data, nil))
	}
}
