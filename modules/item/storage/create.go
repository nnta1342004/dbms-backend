package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) Create(ctx context.Context, data *itemmodel.Item) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
