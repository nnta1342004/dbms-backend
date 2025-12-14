package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) SoftDelete(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(itemmodel.Item{}.TableName()).Where(conditions).Update("status", itemmodel.StatusDeleted)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
