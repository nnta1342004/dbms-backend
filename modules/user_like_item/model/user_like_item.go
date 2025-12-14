package userlikeitemmodel

import (
	"errors"
	"hareta/appCommon"
)

const EntityName = "UserLikeItem"

type UserLikeItem struct {
	UserId  int64 `json:"-" gorm:"column:user_id;primaryKey"`
	GroupId int64 `json:"-" gorm:"column:group_id;primaryKey"`
}

func (UserLikeItem) TableName() string {
	return "user_like_item"
}

type UserLikeItemCreate struct {
	GroupId string `json:"group_id" binding:"required"`
}
type UserLikeItemDelete struct {
	GroupId string `json:"group_id" binding:"required"`
}

var (
	ErrUserHasLikedItem = appCommon.NewCustomError(
		errors.New("user has liked this item"),
		"User has liked this item",
		"ErrUserHasLikedItem",
	)
	ErrUserHasNotLikedItem = appCommon.NewCustomError(
		errors.New("user has not liked this item"),
		"User has not liked this item",
		"ErrUserHasNotLikedItem",
	)
)
