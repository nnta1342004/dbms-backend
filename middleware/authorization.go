package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/tokenprovider"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
	"strings"
)

type authStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *appCommon.AppError {
	return appCommon.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequiredAuth(sc goservice.ServiceContext, authStore authStore) func(c *gin.Context) {
	tokenProvider := sc.MustGet(appCommon.PluginJwt).(tokenprovider.Provider)
	rdc := sc.MustGet(appCommon.PluginRedis).(*redis.Client)
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		id, err := appCommon.FromBase58(payload.UserId)
		res, err := rdc.Exists(c.Request.Context(), fmt.Sprintf("%s_%d_%s", appCommon.UserSessionId, id.GetLocalID(), payload.SessionID)).Result()
		if err != nil {
			panic(appCommon.ErrInternal(err))
		}
		if res != 1 {
			panic(appCommon.ErrNoPermission(errors.New("your session has been expired")))
		}

		user, err := authStore.FindDataWithCondition(c.Request.Context(), map[string]interface{}{"id": id.GetLocalID()}, "Avatar")

		if err != nil {
			panic(err)
		}
		if user.Status != usermodel.StatusVerified {
			panic(usermodel.ErrEmailNotVerified)
		}
		c.Set(appCommon.CurrentUser, user)
		c.Next()
	}
}
