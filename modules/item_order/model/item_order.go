package itemordermodel

import (
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
	ordermodel "hareta/modules/order/model"
)

const EntityName = "ItemOrder"

type ItemOrder struct {
	appCommon.SQLModel
	ItemId   int64             `json:"-" gorm:"column:item_id;index:idx_item_order,unique"`  // Part of composite unique index
	OrderId  int64             `json:"-" gorm:"column:order_id;index:idx_item_order,unique"` // Part of composite unique index
	Quantity int64             `json:"quantity" gorm:"column:quantity"`
	Item     *itemmodel.Item   `json:"item" gorm:"foreignKey:ItemId;references:Id"`
	Order    *ordermodel.Order `json:"-" gorm:"foreignKey:OrderId;references:Id"`
}

func (ItemOrder) TableName() string {
	return "item_order"
}

func (s *ItemOrder) Mask(isAdmin bool) {
	s.GenUID(appCommon.DbTypeItemOrder)
}

type ItemOrderList struct {
	OrderId string `form:"order_id" json:"order_id" binding:"required"`
}
