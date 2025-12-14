package mailer

import (
	"context"
	"hareta/appCommon"
	"hareta/components/asyncjob"
	"net/smtp"
)

type Mail struct {
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Receiver string `json:"receiver"`
}

type MailSender struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewMailSender(Host, Port, Username, Password string) *MailSender {
	return &MailSender{
		Host:     Host,
		Port:     Port,
		Address:  Host + ":" + Port,
		Username: Username,
		Password: Password,
	}
}

func (u *MailSender) SendMail(mail Mail) error {
	auth := smtp.PlainAuth("", u.Username, u.Password, u.Host)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte("Subject:" + mail.Subject + "\n" + mime + mail.Body)

	job := asyncjob.NewJob(func(ctx context.Context) error {
		return smtp.SendMail(u.Address, auth, u.Username, []string{mail.Receiver}, msg)
	})
	err := job.ExecuteWithRetry(context.Background())
	if err != nil {
		return appCommon.NewCustomError(err, "cannot send mail", "ErrCannotSendMail")
	}
	return nil
}
