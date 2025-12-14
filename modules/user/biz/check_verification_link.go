package userbiz

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
	"strconv"
)

type checkVerificationStore interface {
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}
type checkLinkBiz struct {
	store checkVerificationStore
}

func NewCheckLinkBiz(store checkVerificationStore) *checkLinkBiz {
	return &checkLinkBiz{store: store}
}

func (biz *checkLinkBiz) CheckLink(ctx context.Context, rdc *redis.Client, link string) error {
	s, err := rdc.Get(ctx, fmt.Sprintf("%s_%s", appCommon.UserVerification, link)).Result()
	if err != nil {
		if err == redis.Nil {
			return usermodel.ErrLinkIsInvalid
		}
		return appCommon.ErrInternal(err)
	}
	userId, err := strconv.Atoi(s)
	if err != nil {
		return appCommon.ErrInternal(err)
	}
	if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{"id": userId}, map[string]interface{}{"status": usermodel.StatusVerified}); err != nil {
		return appCommon.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	rdc.Del(ctx, fmt.Sprintf("%s_%s", appCommon.UserVerification, link))
	rdc.Del(ctx, fmt.Sprintf("%s_%d", appCommon.UserVerificationTime, userId))

	return nil
}
