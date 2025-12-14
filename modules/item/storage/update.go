package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) Update(ctx context.Context, conditions map[string]interface{}, data *itemmodel.ItemUpdate) error {
	db := s.db.Table(itemmodel.Item{}.TableName()).Where(conditions).Updates(data)
	db.Update("status", data.Status)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
