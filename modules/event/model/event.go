package eventmodel

import (
	"hareta/appCommon"
	eventitemmodel "hareta/modules/event_item/model"
	"time"
)

const EntityName = "Event"

type Event struct {
	appCommon.SQLModel `json:",inline"`
	DateStart          time.Time                  `json:"date_start" gorm:"column:date_start"`
	DateEnd            time.Time                  `json:"date_end" gorm:"column:date_end"`
	OverallContent     string                     `json:"overall_content" gorm:"column:overall_content"`
	DetailContent      string                     `json:"detail_content" gorm:"column:detail_content"`
	Discount           int64                      `json:"discount" gorm:"column:discount"`
	Avatar             string                     `json:"avatar" gorm:"column:avatar"`
	Items              []eventitemmodel.EventItem `json:"items" gorm:"foreignKey:EventId;references:Id"`
}

func (Event) TableName() string {
	return "event"
}

func (s *Event) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeEvent)
}

type EventCreate struct {
	DateStart      int64  `json:"date_start" binding:"required"`
	DateEnd        int64  `json:"date_end" binding:"required"`
	OverallContent string `json:"overall_content"`
	DetailContent  string `json:"detail_content"`
	Discount       int64  `json:"discount"`
	Avatar         string `json:"avatar"`
}

type EventUpdate struct {
	Id             string  `json:"id" binding:"required"`
	DateStart      *int64  `json:"date_start"`
	DateEnd        *int64  `json:"date_end"`
	OverallContent *string `json:"overall_content"`
	DetailContent  *string `json:"detail_content"`
	Discount       *int64  `json:"discount"`
	Avatar         *string `json:"avatar"`
}

type EventDelete struct {
	Id string `json:"id" binding:"required"`
}
