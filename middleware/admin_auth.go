package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
)

func RequiredAdminAuth(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(appCommon.CurrentUser).(*usermodel.User)
		if user.Role == "admin" {
			c.Next()
			return
		}
		panic(appCommon.ErrNoPermission(errors.New("you dont have permission")))
	}
}
