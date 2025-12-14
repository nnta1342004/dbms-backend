package passwordrecoverybiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
	"time"
)

type findStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*passwordrecoverymodel.PasswordRecovery, error)
}
type findBiz struct {
	logger logger.Logger
	store  findStore
}

func NewFindBiz(store findStore) *findBiz {
	return &findBiz{store: store, logger: logger.GetCurrent().GetLogger("PasswordRecoveryFindBiz")}
}

func (biz *findBiz) Find(ctx context.Context, data *passwordrecoverymodel.PasswordRecoveryFind) (*passwordrecoverymodel.PasswordRecovery, error) {
	item, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"slug": data.Slug,
	})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		if err == appCommon.RecordNotFound {
			return nil, appCommon.ErrEntityNotFound(passwordrecoverymodel.EntityName, err)
		}
		return nil, appCommon.ErrCannotGetEntity(passwordrecoverymodel.EntityName, err)
	}
	diff := time.Now().Sub(*item.UpdatedAt)
	if diff.Minutes() > 5 {
		return nil, passwordrecoverymodel.ErrLinkHasBeenExpired
	}
	item.Mask(false)
	return item, nil
}
