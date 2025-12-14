package usergin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
	"net/http"
)

func GetMe(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		user.Mask(false)
		if user.Avatar != nil {
			user.Avatar.Mask(false)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(user))
	}
}
