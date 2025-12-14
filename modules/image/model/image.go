package imagemodel

import (
	"errors"
	"hareta/appCommon"
	"time"
)

const EntityName = "image"

type Image struct {
	appCommon.SQLModel `json:",inline"`
	URL                string `json:"url" binding:"required" gorm:"column:url;type:varchar(200)"`
	FileName           string `json:"file_name" gorm:"column:file_name;type:varchar(200)"`
	Width              int    `json:"width" gorm:"column:width"`
	Height             int    `json:"height" gorm:"column:height"`
	CloudName          string `json:"cloud_name" gorm:"column:cloud_name;type:varchar(100)"`
	Extension          string `json:"extension" gorm:"column:extension;type:varchar(100)"`
}

func (Image) TableName() string {
	return "image"
}
func (s *Image) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeImage)
}

type ImageDelete struct {
	ImageId string `json:"image_id" binding:"required"`
}

type ImageList struct {
	TimeFrom *int64 `form:"time_from"`
	TimeTo   *int64 `form:"time_to"`
}
type ImageFilter struct {
	TimeFrom time.Time
	TimeTo   time.Time
}

func (s *ImageList) FulFill() {
	if s.TimeTo == nil {
		s.TimeTo = new(int64)
		*s.TimeTo = time.Now().Unix()
	}
	if s.TimeFrom == nil {
		s.TimeFrom = new(int64)
		*s.TimeFrom = 0
	}
}

var (
	ErrCannotFindImage   = appCommon.NewCustomError(errors.New("Cannot find the image"), "Cannot find the image", "ErrCannotFindImage")
	ErrCannotUploadImage = appCommon.NewCustomError(
		errors.New("Cannot upload the image"),
		"Cannot upload the image",
		"ErrCannotUploadImage",
	)
	ErrInvalidImageFormat = appCommon.NewCustomError(nil, "Invalid image format", "ErrInvalidImageFormat")
)
