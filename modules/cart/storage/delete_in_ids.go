package cartstorage

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
)

func (s *sqlStore) DeleteInIds(ctx context.Context, ids []int64) error {
	db := s.db.Table(cartmodel.Cart{}.TableName()).Where("id IN ?", ids)
	db = db.Delete(&cartmodel.Cart{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
