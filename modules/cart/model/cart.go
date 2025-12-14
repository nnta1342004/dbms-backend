package cartmodel

import (
	"errors"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
	usermodel "hareta/modules/user/model"
)

const EntityName = "cart"
const (
	StatusActive = iota
	StatusDeleted
)

type Cart struct {
	appCommon.SQLModel `json:",inline"`
	UserId             int64           `json:"-" gorm:"column:user_id"`
	ItemId             int64           `json:"-" gorm:"column:item_id"`
	Quantity           int64           `json:"quantity" gorm:"column:quantity"`
	User               *usermodel.User `json:"-" gorm:"foreignKey:UserId;references:Id;"`
	Item               *itemmodel.Item `json:"item" gorm:"foreignKey:ItemId;references:Id;"`
}

type CartCreate struct {
	ItemId   string `json:"item_id" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required"`
}
type CartDelete struct {
	Id string `json:"id" binding:"required"`
}
type CartUpdate struct {
	Id       string `json:"id" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required"`
}

func (Cart) TableName() string {
	return "cart"
}

func (s *Cart) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeCart)
}

var (
	ErrQuantityExceed = appCommon.NewCustomError(
		errors.New("your quantity exceed the item store"),
		"your quantity exceed the item store",
		"ErrQuantityExceed",
	)
)
