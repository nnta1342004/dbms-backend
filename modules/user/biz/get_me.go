package userbiz

import (
	"context"
	usermodel "hareta/modules/user/model"
)

type meStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type meBiz struct {
	store meStore
}

func NewMeBiz(store meStore) *meBiz {
	return &meBiz{store: store}
}

func (biz *meBiz) GetMe(ctx context.Context, user *usermodel.User) (*usermodel.User, error) {
	user.Mask(false)
	return user, nil
}
