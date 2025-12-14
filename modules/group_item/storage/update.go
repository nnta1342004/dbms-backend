package groupitemstorage

import (
	"context"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(groupitemmodel.GroupItem{}.TableName()).Where("status != ?", itemmodel.StatusDeleted).Where(conditions).Updates(data)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
