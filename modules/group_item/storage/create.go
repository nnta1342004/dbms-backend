package groupitemstorage

import (
	"context"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
)

func (s *sqlStore) Create(ctx context.Context, data *groupitemmodel.GroupItem) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
