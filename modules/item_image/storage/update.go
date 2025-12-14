package itemimagestorage

import (
	"context"
	"hareta/appCommon"
	itemimagemodel "hareta/modules/item_image/model"
)

func (s *sqlStore) UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(itemimagemodel.ItemImage{}.TableName()).Where(conditions).Updates(data)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
