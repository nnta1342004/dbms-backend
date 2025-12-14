package ordermodel

import (
	"errors"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
)

const (
	StatusPreparing = iota
	StatusTransferring
	StatusDelivered
	StatusRejected
)

const EntityName = "Order"

type Order struct {
	appCommon.SQLModel `json:",inline"`
	Address            string          `json:"address" gorm:"column:address;type:varchar(200)"`
	Name               string          `json:"name" gorm:"column:name;type:varchar(100)"`
	Email              string          `json:"email" gorm:"column:email;type:varchar(100)"`
	Phone              string          `json:"phone" gorm:"column:phone;type:varchar(30)"`
	UserId             int64           `json:"-" gorm:"column:user_id"`
	Total              int64           `json:"total" gorm:"column:total"`
	User               *usermodel.User `json:"-" gorm:"foreignKey:UserId;references:Id"`
}

func (s *Order) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeOrder)
}

func (Order) TableName() string {
	return "order"
}

type ItemOrder struct {
	Id       string `json:"id" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required"`
}
type OrderCreateWithoutLogin struct {
	Item    []ItemOrder `json:"item" binding:"required"`
	Address string      `json:"address" binding:"required"`
	Email   string      `json:"email" binding:"required"`
	Phone   string      `json:"phone" binding:"required"`
	Name    string      `json:"name" binding:"required"`
}
type OrderCreate struct {
	Id      []string `json:"id" binding:"required"`
	Address string   `json:"address" binding:"required"`
	Email   string   `json:"email" binding:"required"`
	Phone   string   `json:"phone" binding:"required"`
	Name    string   `json:"name" binding:"required"`
}
type OrderUpdate struct {
	Id     string `json:"id" binding:"required"`
	Status int    `json:"status"`
}
type OrderListAdmin struct {
	Status *int    `form:"status"`
	Email  *string `form:"email"`
}
type OrderFind struct {
	OrderId string `json:"order_id" binding:"required" form:"order_id"`
}

var (
	ErrQuantityLimitExceed = appCommon.NewCustomError(
		errors.New("quantity limit exceed"),
		"Quantity limit exceed",
		"ErrQuantityLimitExceed",
	)
)
