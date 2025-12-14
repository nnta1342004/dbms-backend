package cartstorage

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
)

func (s *sqlStore) Create(ctx context.Context, data *cartmodel.Cart) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
