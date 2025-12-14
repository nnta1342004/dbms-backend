package passwordrecoverybiz

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
	usermodel "hareta/modules/user/model"
	"hareta/plugin/pubsub"
	"time"
)

type createStore interface {
	Create(ctx context.Context, user *passwordrecoverymodel.PasswordRecovery) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*passwordrecoverymodel.PasswordRecovery, error)
	UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error
}
type createUserStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type createBiz struct {
	store     createStore
	userStore createUserStore
	rabbitMQ  pubsub.PubSub
	logger    logger.Logger
}

func NewCreateBiz(store createStore, userStore createUserStore, rabbitMQ pubsub.PubSub) *createBiz {
	return &createBiz{
		rabbitMQ:  rabbitMQ,
		store:     store,
		userStore: userStore,
		logger:    logger.GetCurrent().GetLogger("UserForgotPasswordCreateBiz"),
	}
}
func (biz *createBiz) Create(ctx context.Context, data *passwordrecoverymodel.UserRecreatePassword) (*passwordrecoverymodel.PasswordRecovery, error) {

	_, err := biz.userStore.FindDataWithCondition(ctx, map[string]interface{}{
		"email": data.Email,
	})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		if err == appCommon.RecordNotFound {
			return nil, appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return nil, appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	item, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{
		"email": data.Email,
	})
	if err != nil {
		if err != appCommon.RecordNotFound {
			biz.logger.WithSrc().Errorln(err)
			return nil, appCommon.ErrCannotGetEntity(passwordrecoverymodel.EntityName, err)
		}
	}

	if err == appCommon.RecordNotFound {
		createData := &passwordrecoverymodel.PasswordRecovery{
			Email: data.Email,
			Slug:  appCommon.GenSalt(20),
		}
		if err := biz.store.Create(ctx, createData); err != nil {
			biz.logger.WithSrc().Errorln(err)
		}
		createData.Mask(false)
		item = createData
	} else {
		diff := time.Now().Sub(*item.UpdatedAt)
		if diff.Seconds() <= 60 {
			return nil, usermodel.ErrWaitingForANewLink
		}
		item.Slug = appCommon.GenSalt(20)
		if err := biz.store.UpdateWithCondition(ctx, map[string]interface{}{
			"email": data.Email,
		}, map[string]interface{}{
			"slug": item.Slug,
		}); err != nil {
			biz.logger.WithSrc().Errorln(err)
			return nil, appCommon.ErrCannotUpdateEntity(passwordrecoverymodel.EntityName, err)
		}
	}

	dataSend := passwordrecoverymodel.RecoveryPasswordRabbitMQ{
		Email: item.Email,
		Slug:  item.Slug,
	}
	message, err := appCommon.StructToMap(dataSend)
	if err != nil {
		return nil, appCommon.ErrInternal(err)
	}
	if err := biz.rabbitMQ.Publish(ctx, appCommon.TopicSendMailRecoveryPassword, pubsub.NewMessage(message)); err != nil {
		return nil, appCommon.ErrInternal(err)
	}
	item.Mask(false)
	return item, nil
}
