package passwordrecoverymodel

import (
	"errors"
	"hareta/appCommon"
)

const EntityName = "PasswordRecovery"

type PasswordRecovery struct {
	appCommon.SQLModel
	Email string `json:"email" gorm:"column:email;type:varchar(100)"`
	Slug  string `json:"slug" gorm:"column:slug;type:varchar(100)"`
}

func (PasswordRecovery) TableName() string {
	return "password_recovery"
}

func (s *PasswordRecovery) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeUserForgotPassword)
}

type UserRecreatePassword struct {
	Email string `json:"email" binding:"required"`
}
type PasswordRecoveryFind struct {
	Slug string `json:"slug" form:"slug" binding:"required"`
}
type RecoverData struct {
	Password string `json:"password" binding:"required"`
	Slug     string `json:"slug" binding:"required"`
}

type RecoveryPasswordRabbitMQ struct {
	Email string `json:"email" mapstructure:"email"`
	Slug  string `json:"slug" mapstructure:"slug"`
}

var (
	ErrLinkHasBeenExpired = appCommon.NewCustomError(
		errors.New("your link has been expired"),
		"your link has been expired",
		"ErrLinkHasBeenExpired",
	)
)
