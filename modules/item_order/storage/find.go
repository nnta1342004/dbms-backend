package itemorderstorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	itemordermodel "hareta/modules/item_order/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*itemordermodel.ItemOrder, error) {
	db := s.db.Table(itemordermodel.ItemOrder{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data itemordermodel.ItemOrder
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
