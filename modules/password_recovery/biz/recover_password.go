package passwordrecoverybiz

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
	usermodel "hareta/modules/user/model"
	"time"
)

type recoverStore interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*passwordrecoverymodel.PasswordRecovery, error)
}
type recoverUserStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}

type recoverBiz struct {
	store     recoverStore
	userStore recoverUserStore
	logger    logger.Logger
}

func NewRecoverBiz(store recoverStore, userStore recoverUserStore) *recoverBiz {
	return &recoverBiz{
		store:     store,
		userStore: userStore,
		logger:    logger.GetCurrent().GetLogger("RecoverPasswordBiz"),
	}
}

func (biz *recoverBiz) Recover(ctx context.Context, data *passwordrecoverymodel.RecoverData) error {
	item, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"slug": data.Slug,
	})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(passwordrecoverymodel.EntityName, err)
		}
		return appCommon.ErrCannotGetEntity(passwordrecoverymodel.EntityName, err)
	}

	if time.Now().Sub(*item.UpdatedAt) > time.Minute*5 {
		return passwordrecoverymodel.ErrLinkHasBeenExpired
	}

	user, err := biz.userStore.FindDataWithCondition(ctx, map[string]interface{}{
		"email": item.Email,
	})

	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	hash := sha256.New()
	hash.Write([]byte(data.Password + user.Salt))
	newPassword := fmt.Sprintf("%x", hash.Sum(nil))

	if err := biz.userStore.UpdateWithCondition(ctx, map[string]interface{}{
		"id": user.Id,
	}, map[string]interface{}{
		"password": newPassword,
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	if err := biz.store.DeleteWithCondition(ctx, map[string]interface{}{
		"id": item.Id,
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(passwordrecoverymodel.EntityName, err)
	}
	return nil
}
