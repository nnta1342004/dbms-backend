package itemimagestorage

import (
	"context"
	"hareta/appCommon"
	itemimagemodel "hareta/modules/item_image/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(itemimagemodel.ItemImage{}.TableName()).Where(conditions)
	db = db.Delete(&itemimagemodel.ItemImage{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
