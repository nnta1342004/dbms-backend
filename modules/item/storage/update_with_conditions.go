package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) UpdateWithCondition(
	ctx context.Context,
	conditions map[string]interface{},
	data map[string]interface{},
) error {
	db := s.db.Table(itemmodel.Item{}.TableName()).Where("status != ?", itemmodel.StatusDeleted).Where(conditions).Updates(data)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
