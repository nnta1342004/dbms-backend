package eventitemmodel

import (
	itemmodel "hareta/modules/item/model"
)

const EntityName = "EventItem"

type EventItem struct {
	EventId int64           `json:"-" gorm:"column:event_id;primaryKey"`
	ItemId  int64           `json:"-" gorm:"column:item_id;unique;primaryKey"`
	Item    *itemmodel.Item `json:"item" gorm:"foreignKey:ItemId;references:Id"`
}

func (EventItem) TableName() string {
	return "event_item"
}

type EventItemCreate struct {
	EventId string `json:"event_id" binding:"required"`
	ItemId  string `json:"item_id" binding:"required"`
}

type EventItemDelete struct {
	EventId string `json:"event_id" binding:"required"`
	ItemId  string `json:"item_id" binding:"required"`
}
