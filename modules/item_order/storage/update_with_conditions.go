package itemorderstorage

import (
	"context"
	"hareta/appCommon"
	itemordermodel "hareta/modules/item_order/model"
)

func (s *sqlStore) UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(itemordermodel.ItemOrder{}.TableName()).Where(conditions).Updates(data)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
