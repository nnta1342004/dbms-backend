package userbiz

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/leductoan3082004/go-sdk/plugin/tokenprovider"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
)

type loginStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBiz struct {
	store loginStore
}

func NewLoginBiz(store loginStore) *loginBiz {
	return &loginBiz{store: store}
}
func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin, tokenProvider tokenprovider.Provider, rdc *redis.Client) (*tokenprovider.Token, error) {
	item, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return nil, appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return nil, appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	hash := sha256.New()
	hash.Write([]byte(data.Password + item.Salt))
	if item.Password != fmt.Sprintf("%x", hash.Sum(nil)) {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	if item.Status != usermodel.StatusVerified {
		return nil, usermodel.ErrEmailNotVerified
	}

	item.Mask(false)
	sessionId := appCommon.GenSalt(20)
	payload := tokenprovider.TokenPayload{
		UserId:    item.FakeId.String(),
		Role:      item.Role,
		SessionID: sessionId,
	}
	_, err = rdc.Set(ctx, fmt.Sprintf("%s_%d_%s", appCommon.UserSessionId, item.Id, sessionId), "1", 0).Result()
	token, err := tokenProvider.Generate(payload, appCommon.ExpiryAccessToken)
	if err != nil {
		return nil, tokenprovider.ErrEncodingToken
	}
	return token, err
}
