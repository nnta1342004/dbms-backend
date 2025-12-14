package userbiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
)

type updateStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}
type updateBiz struct {
	store  updateStore
	logger logger.Logger
}

func NewUpdateBiz(store updateStore) *updateBiz {
	return &updateBiz{store: store, logger: logger.GetCurrent().GetLogger("UserUpdateBiz")}
}

func (biz *updateBiz) Update(ctx context.Context, user *usermodel.User, update *usermodel.UserUpdate) error {
	updateMap := make(map[string]interface{})
	if update.Name != nil {
		if len(*update.Name) > 0 {
			updateMap["name"] = *update.Name
		}
	}
	if update.Phone != nil {
		updateMap["phone"] = *update.Phone
	}

	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{
		"id": user.Id,
	}, updateMap); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
