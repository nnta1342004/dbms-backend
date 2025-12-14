package usermodel

import (
	"errors"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
)

const (
	EntityName = "user"
)
const (
	StatusNotVerified = iota
	StatusVerified
)

type User struct {
	appCommon.SQLModel `json:",inline"`
	Name               string            `json:"name" gorm:"column:name;type:varchar(100)"`
	Email              string            `json:"email" gorm:"column:email;type:varchar(100)"`
	Password           string            `json:"-" gorm:"column:password;type:varchar(100)"`
	Salt               string            `json:"-" gorm:"column:salt;type:varchar(100)"`
	Role               string            `json:"role" gorm:"column:role;type:varchar(50)"`
	Phone              string            `json:"phone" gorm:"column:phone;type:varchar(30)"`
	AvatarId           int64             `json:"-" gorm:"column:avatar_id;default:null"`
	Avatar             *imagemodel.Image `json:"avatar" gorm:"foreignKey:AvatarId;references:Id"`
}

type UserUpdate struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}
type UserCreate struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}
type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserLink struct {
	Email string `json:"email" binding:"required"`
}
type UserChangePassword struct {
	OldPassword        string `json:"old_password" binding:"required"`
	NewPassword        string `json:"new_password" binding:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" binding:"required"`
}

func (User) TableName() string {
	return "user"
}

func (s *User) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeUser)
}

var (
	ErrUsernameExisted = appCommon.NewCustomError(
		errors.New("username Existed"),
		"Username Existed",
		"ErrUsernameExisted",
	)
	ErrEmailExisted = appCommon.NewCustomError(
		errors.New("email Existed"),
		"Email Existed",
		"ErrEmailExisted",
	)
	ErrEmailNotVerified = appCommon.NewCustomError(
		errors.New("email is not verified"),
		"Email is not verified",
		"ErrEmailNotVerified",
	)
	ErrUsernameOrPasswordInvalid = appCommon.NewCustomError(
		errors.New("username or password is invalid"),
		"username or password is invalid",
		"ErrUsernameOrPasswordInvalid",
	)
	ErrEmailVerified = appCommon.NewCustomError(
		errors.New("email has been verified"),
		"Email has been verified",
		"ErrEmailVerified",
	)
	ErrWaitingForANewLink = appCommon.NewCustomError(
		errors.New("you have to wait for 1 minute to resend the link"),
		"You have to wait for 1 minute to resend the link",
		"ErrWaitingForANewLink",
	)
	ErrLinkIsInvalid = appCommon.NewCustomError(
		errors.New("your link is invalid"),
		"Your link is invalid",
		"ErrLinkIsInvalid",
	)
	ErrNewPasswordIsInvalid = appCommon.NewCustomError(
		errors.New("your new password is not matched"),
		"Your new password is not matched",
		"ErrNewPasswordIsInvalid",
	)
	ErrOldPasswordIsInvalid = appCommon.NewCustomError(
		errors.New("your old password is incorrect"),
		"Your old password is incorrect",
		"ErrOldPasswordIsInvalid",
	)
)
