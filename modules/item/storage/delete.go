package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(itemmodel.Item{}.TableName()).Where(conditions)
	db = db.Delete(&itemmodel.Item{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
