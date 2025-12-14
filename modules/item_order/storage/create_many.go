package itemorderstorage

import (
	"context"
	"hareta/appCommon"
	itemordermodel "hareta/modules/item_order/model"
)

func (s *sqlStore) CreateMany(ctx context.Context, data []itemordermodel.ItemOrder) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
