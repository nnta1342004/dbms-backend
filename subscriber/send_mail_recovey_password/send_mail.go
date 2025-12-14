package passwordrecoverysub

import (
	"context"
	"fmt"
	goservice "github.com/leductoan3082004/go-sdk"
	mailplugin "github.com/leductoan3082004/go-sdk/plugin/mailer"
	mailengine "github.com/leductoan3082004/go-sdk/plugin/mailer/mail"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"hareta/appCommon"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
	"hareta/plugin/pubsub"
	"io/ioutil"
	"os"
	"strings"
)

func SendMailRecoveryPassword(sc goservice.ServiceContext, topic string, numWorkers int) appCommon.SubJob {
	return appCommon.SubJob{
		Title:     "send mail recovery password",
		NumWorker: numWorkers,
		Topic:     topic,
		Hld: func(ctx context.Context, message *pubsub.Message, workerId int) error {
			logrus.Infoln("recovery subscriber")
			var data passwordrecoverymodel.RecoveryPasswordRabbitMQ
			if err := mapstructure.Decode(message.Data(), &data); err != nil {
				return appCommon.ErrInternal(err)
			}
			mail := sc.MustGet(appCommon.PluginMailer).(mailplugin.MailEngine)

			htmlFilePath := "./passwordRecovery.html"
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
			body = strings.Replace(
				body, "{{.link}}", fmt.Sprintf("https://www.haretaworkshop.com/password-recovery/%s", data.Slug), -1,
			)
			body = strings.Replace(body, "{{.email}}", data.Email, -1)
			mailContent := mailengine.Mail{
				Subject:  "Recovery password for hareta",
				Body:     body,
				Receiver: data.Email,
			}
			if err := mail.SendMail(mailContent); err != nil {
				return appCommon.ErrInternal(err)
			}
			return nil
		},
	}
}
