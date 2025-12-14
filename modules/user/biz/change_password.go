package userbiz

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/go-redis/redis/v8"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
)

type changePasswordStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type changePasswordBiz struct {
	store changePasswordStore
}

func NewChangePasswordBiz(store changePasswordStore) *changePasswordBiz {
	return &changePasswordBiz{store: store}
}

func (biz *changePasswordBiz) ChangePassword(ctx context.Context, user *usermodel.User, data *usermodel.UserChangePassword, rdc *redis.Client) error {
	if data.NewPassword != data.ConfirmNewPassword {
		return usermodel.ErrNewPasswordIsInvalid
	}
	item, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": user.Id})

	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	hash := sha256.New()
	hash.Write([]byte(data.OldPassword + item.Salt))
	prevPassword := fmt.Sprintf("%x", hash.Sum(nil))

	if prevPassword != item.Password {
		return usermodel.ErrOldPasswordIsInvalid
	}
	hash = sha256.New()
	hash.Write([]byte(data.NewPassword + item.Salt))
	newPassword := fmt.Sprintf("%x", hash.Sum(nil))

	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{"id": user.Id}, map[string]interface{}{
		"password": newPassword,
	}); err != nil {
		return appCommon.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	keys, _ := rdc.Keys(ctx, fmt.Sprintf("%s_%d*", appCommon.UserSessionId, item.Id)).Result()
	for i := range keys {
		rdc.Del(ctx, keys[i])
	}
	return nil
}
