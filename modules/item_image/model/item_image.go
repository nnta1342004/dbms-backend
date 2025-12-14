package itemimagemodel

import (
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
	itemmodel "hareta/modules/item/model"
)

type ItemImage struct {
	appCommon.SQLModel
	ItemId  int64             `json:"-" gorm:"column:item_id"`
	ImageId int64             `json:"-" gorm:"column:image_id"`
	Color   string            `json:"color" gorm:"column:color;type:varchar(30)"`
	Item    *itemmodel.Item   `json:"-" gorm:"foreignKey:ItemId;references:Id"`
	Image   *imagemodel.Image `json:"image" gorm:"foreignKey:ImageId;references:Id"`
}

type ItemCreate struct {
	ItemId string `json:"item_id" form:"item_id" binding:"required"`
	Color  string `json:"color" form:"color" binding:"required"`
}
type ItemDelete struct {
	Id string `json:"id" binding:"required"`
}
type ItemList struct {
	Id               string `json:"id" form:"id" binding:"required"`
	appCommon.Paging `json:",inline"`
}
type ItemUpdate struct {
	Id    string `json:"id" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (s *ItemImage) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeItemImage)
}

const EntityName = "ItemImage"

func (ItemImage) TableName() string {
	return "item_image"
}
