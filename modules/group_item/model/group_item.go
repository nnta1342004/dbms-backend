package groupitemmodel

import (
	"hareta/appCommon"
)

const EntityName = "GroupItem"

type GroupItem struct {
	appCommon.SQLModel
	Name string `json:"name" gorm:"column:name;type:varchar(100)"`
}

func (GroupItem) TableName() string {
	return "group_item"
}

func (s *GroupItem) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeGroupItem)
}

type GroupCreate struct {
	Name string `json:"name" binding:"required"`
}
type GroupDelete struct {
	Id string `json:"id" binding:"required"`
}
type GroupUpdate struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type GroupPriceUpdate struct {
	Price         *int64 `json:"price"`
	OriginalPrice *int64 `json:"original_price"`
}
