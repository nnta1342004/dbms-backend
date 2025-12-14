package userbiz

import (
	"context"
	"crypto/sha256"
	"fmt"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
)

type userStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (
		*usermodel.User, error,
	)
	Create(ctx context.Context, user *usermodel.User) error
}

type userBiz struct {
	store userStore
}

func NewUserBiz(store userStore) *userBiz {
	return &userBiz{store: store}
}
func (biz *userBiz) Register(ctx context.Context, user *usermodel.UserCreate) error {
	_, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"email": user.Email})
	if err != nil && err != appCommon.RecordNotFound {
		return appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	if err == nil {
		return usermodel.ErrEmailExisted
	}

	salt := appCommon.GenSalt(30)
	hash := sha256.New()
	hash.Write([]byte(user.Password + salt))

	userCreate := usermodel.User{
		Email:    user.Email,
		Password: fmt.Sprintf("%x", hash.Sum(nil)),
		Salt:     salt,
		Role:     "user",
		Phone:    user.Phone,
		AvatarId: 0,
		Name:     user.Name,
	}
	if err := biz.store.Create(ctx, &userCreate); err != nil {
		return appCommon.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
