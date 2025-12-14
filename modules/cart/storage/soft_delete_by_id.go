package cartstorage

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
)

func (s *sqlStore) SoftDeleteByItemId(ctx context.Context, itemId int64) error {
	db := s.db.
		Table(cartmodel.Cart{}.TableName()).
		Where("item_id = ?", itemId).
		Delete(&cartmodel.Cart{})

	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
