package userbiz

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/leductoan3082004/go-sdk/logger"
	mailplugin "github.com/leductoan3082004/go-sdk/plugin/mailer"
	mailengine "github.com/leductoan3082004/go-sdk/plugin/mailer/mail"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type sendLinkStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (
		*usermodel.User, error,
	)
}

type sendLinkBiz struct {
	logger logger.Logger
	store  sendLinkStore
}

func NewSendLinkBiz(store sendLinkStore) *sendLinkBiz {
	return &sendLinkBiz{store: store, logger: logger.GetCurrent().GetLogger("SendLinkBiz")}
}

func (biz *sendLinkBiz) SendLink(
	ctx context.Context, email string, mail mailplugin.MailEngine, rdc *redis.Client,
) error {
	user, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"email": email})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return appCommon.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if user.Status == usermodel.StatusVerified {
		return usermodel.ErrEmailVerified
	}

	res, err := rdc.Exists(ctx, fmt.Sprintf("%s_%d", appCommon.UserVerificationTime, user.Id)).Result()
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}
	if res == 1 {
		return usermodel.ErrWaitingForANewLink
	}

	link := appCommon.GenSalt(15)

	tx := rdc.TxPipeline()
	tx.Set(ctx, fmt.Sprintf("%s_%s", appCommon.UserVerification, link), fmt.Sprintf("%d", user.Id), time.Minute*5)
	tx.Set(ctx, fmt.Sprintf("%s_%d", appCommon.UserVerificationTime, user.Id), "1", time.Minute)
	_, err = tx.Exec(ctx)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrInternal(err)
	}

	htmlFilePath := "./verifyEmail.html"
	htmlFile, err := os.Open(htmlFilePath)
	if err != nil {
		return appCommon.ErrInternal(err)
	}
	defer htmlFile.Close()

	htmlContent, err := ioutil.ReadAll(htmlFile)
	if err != nil {
		return appCommon.ErrInternal(err)
	}
	body := string(htmlContent)
	body = strings.Replace(body, "{{.link}}", fmt.Sprintf("https://www.haretaworkshop.com/verify-email/%s", link), -1)
	body = strings.Replace(body, "{{.email}}", user.Email, -1)

	data := mailengine.Mail{
		Subject:  "Your verification link for hareta",
		Body:     body,
		Receiver: email,
	}
	if err := mail.SendMail(data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return err
	}
	return nil
}
