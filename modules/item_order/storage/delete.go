package itemorderstorage

import (
	"context"
	"hareta/appCommon"
	itemordermodel "hareta/modules/item_order/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(itemordermodel.ItemOrder{}.TableName()).Where(conditions)
	db = db.Delete(&itemordermodel.ItemOrder{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
