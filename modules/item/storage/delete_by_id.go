package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) SoftDeleteById(ctx context.Context, id int64) error {
	db := s.db.
		Table(itemmodel.Item{}.TableName()).
		Where("id = ?", id).
		Update("status", itemmodel.StatusDeleted)

	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
