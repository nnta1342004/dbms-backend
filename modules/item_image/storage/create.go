package itemimagestorage

import (
	"context"
	"hareta/appCommon"
	itemimagemodel "hareta/modules/item_image/model"
)

func (s *sqlStore) Create(ctx context.Context, data *itemimagemodel.ItemImage) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
