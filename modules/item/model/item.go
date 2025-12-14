package itemmodel

import (
	"github.com/mitchellh/mapstructure"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
	imagemodel "hareta/modules/image/model"
	"math"
)

const EntityName = "Item"
const (
	StatusActive = iota
	StatusDeleted
	StatusOutOfStock
)

const (
	TagNothing = iota
	TagTopSeller
	TagSignature
	TagFavorite
)

const (
	CronAutoUpdate = iota
	CronNoUpdate
)

type ItemSystem struct {
	Sold       int64 `json:"sold" gorm:"column:sold"`
	LikeCount  int64 `json:"like_count" gorm:"column:like_count"`
	Tag        int   `json:"tag" gorm:"column:tag"`
	CronStatus int   `json:"cron_status" gorm:"column:cron_status"`
}

func (s *ItemSystem) HideSensitiveData() {
	s.Sold = 0
	s.CronStatus = 0
}

type ItemAttr struct {
	Name          string `json:"name" gorm:"column:name"`
	ProductLine   string `json:"product_line" gorm:"column:product_line"`
	Category      string `json:"category" gorm:"column:category"`
	Quantity      int64  `json:"quantity" gorm:"column:quantity"`
	Discount      int64  `json:"discount" gorm:"column:discount"`
	Collection    string `json:"collection" gorm:"column:collection"`
	Type          string `json:"type" gorm:"column:type"`
	Price         int64  `json:"price" gorm:"column:price"`
	OriginalPrice int64  `json:"original_price" gorm:"column:original_price"`
	Color         string `json:"color" gorm:"column:color"`
}
type Item struct {
	appCommon.SQLModel `json:",inline"`
	ItemAttr           `json:",inline"`
	ItemSystem         `json:",inline"`
	GroupId            int64                     `json:"-" gorm:"column:group_id"`
	Group              *groupitemmodel.GroupItem `json:"group,omitempty" gorm:"foreignKey:GroupId"`
	Description        string                    `json:"description" gorm:"column:description"`
	AvatarId           int64                     `json:"-" gorm:"column:avatar_id"`
	Default            bool                      `json:"default" gorm:"column:default"`
	Avatar             *imagemodel.Image         `json:"avatar,omitempty" gorm:"foreignKey:AvatarId"`
}
type SimpleItem struct {
	appCommon.SQLModel `json:",inline"`
	ItemAttr           `json:",inline"`
	ItemSystem         `json:",inline"`
	Default            bool                      `json:"default" gorm:"column:default"`
	GroupId            int64                     `json:"-" gorm:"column:group_id"`
	AvatarId           int64                     `json:"-" gorm:"column:avatar_id"`
	Group              *groupitemmodel.GroupItem `json:"group,omitempty" gorm:"foreignKey:GroupId"`
	Avatar             *imagemodel.Image         `json:"avatar"`
}

func (SimpleItem) TableName() string {
	return "item"
}

const ItemTypeName = "ItemType"

var Type = []string{"category", "collection", "type", "product_line"}

func (s *PriceFilter) FulFill() {
	if s.LowerPrice == nil {
		s.LowerPrice = new(int64)
		*s.LowerPrice = 0
	}
	if s.UpperPrice == nil {
		s.UpperPrice = new(int64)
		*s.UpperPrice = math.MaxInt64
	}
}
func (Item) TableName() string {
	return "item"
}
func (ItemUpdate) TableName() string {
	return "item"
}

func (s *Item) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeItem)
}
func (s *SimpleItem) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeItem)
}

type ItemCreate struct {
	GroupId       string `json:"group_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Category      string `json:"category" binding:"required"`
	Discount      int64  `json:"discount,omitempty"`
	Quantity      int64  `json:"quantity" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Collection    string `json:"collection" binding:"required"`
	Type          string `json:"type" binding:"required"`
	OriginalPrice int64  `json:"original_price" binding:"required"`
	Price         int64  `json:"price" binding:"required"`
	ProductLine   string `json:"product_line" binding:"required"`
	Color         string `json:"color" binding:"required"`
}
type PriceFilter struct {
	LowerPrice *int64 `json:"lower_price" form:"lower_price"`
	UpperPrice *int64 `json:"upper_price" form:"upper_price"`
	Desc       *bool  `json:"desc" form:"desc"`
}
type TypeFilter struct {
	Category    *string `json:"category" form:"category"`
	Collection  *string `json:"collection" form:"collection"`
	Type        *string `json:"type" form:"type"`
	ProductLine *string `json:"product_line" form:"product_line"`
	Tag         *int    `json:"tag" form:"tag"`
}
type SearchFilter struct {
	Name *string `json:"name" form:"name"`
}
type ItemList struct {
	PriceFilter
	TypeFilter
	SearchFilter
}

type ItemUpdateAttr struct {
	Status        *int    `json:"status" mapstructure:"status,omitempty"`
	Name          *string `json:"name" mapstructure:"name,omitempty"`
	Category      *string `json:"category" mapstructure:"category,omitempty"`
	Quantity      *int64  `json:"quantity" mapstructure:"quantity,omitempty"`
	Description   *string `json:"description" mapstructure:"description,omitempty"`
	Collection    *string `json:"collection" mapstructure:"collection,omitempty"`
	Type          *string `json:"type" mapstructure:"type,omitempty"`
	Price         *int64  `json:"price" mapstructure:"price,omitempty"`
	OriginalPrice *int64  `json:"original_price" mapstructure:"original_price,omitempty"`
	ProductLine   *string `json:"product_line" mapstructure:"product_line,omitempty"`
	Discount      *int64  `json:"discount" mapstructure:"discount,omitempty"`
}

type ItemUpdateSystem struct {
	CronStatus *int `json:"cron_status" mapstructure:"cron_status,omitempty"`
	Tag        *int `json:"tag" mapstructure:"tag,omitempty"`
}
type ItemUpdateAll struct {
	ItemUpdateSystem `json:",inline" mapstructure:",inline"`
	ItemUpdateAttr   `json:",inline" mapstructure:",inline"`
}

func (s *ItemUpdateAttr) ToMap() (map[string]interface{}, error) {
	mData := make(map[string]interface{})
	if err := mapstructure.Decode(s, &mData); err != nil {
		return nil, appCommon.ErrInternal(err)
	}
	return mData, nil
}

func (s *ItemUpdateSystem) ToMap() (map[string]interface{}, error) {
	mData := make(map[string]interface{})
	if err := mapstructure.Decode(s, &mData); err != nil {
		return nil, appCommon.ErrInternal(err)
	}
	return mData, nil
}
func (s *ItemUpdateAll) ToMap() (map[string]interface{}, error) {
	m1, err := s.ItemUpdateSystem.ToMap()
	if err != nil {
		return nil, err
	}
	m2, err := s.ItemUpdateAttr.ToMap()
	if err != nil {
		return nil, err
	}
	for key, item := range m2 {
		m1[key] = item
	}
	return m1, nil
}

type ItemUpdate struct {
	Id string `json:"id" gorm:"-" binding:"required"`
	ItemUpdateAll
	AvatarId     int64   `json:"-" gorm:"avatar_id"`
	AvatarFakeId *string `json:"avatar_id" gorm:"-"`
}

type ItemAvtUpdate struct {
	Id      string
	ImageId int64
}
type ItemFind struct {
	Id string `json:"id" binding:"required" form:"id"`
}
type ItemDelete struct {
	Id string `json:"id" binding:"required"`
}
type ItemCronUpdate struct {
	Id         string `json:"id" binding:"required"`
	Tag        *int   `json:"tag"`
	CronStatus *int   `json:"cron_status"`
}
type ItemTypeList struct {
	Query string `json:"query" form:"query"`
	Field string `json:"field" form:"field" binding:"required"`
}

type ItemMakeDefault struct {
	Id string `json:"id" binding:"required"`
}

type ItemListGroup struct {
	Id string `json:"id" form:"id" binding:"required"`
}
