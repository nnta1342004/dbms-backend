package groupitemstorage

import (
	"context"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(groupitemmodel.GroupItem{}.TableName()).Where(conditions)
	db = db.Delete(&groupitemmodel.GroupItem{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
