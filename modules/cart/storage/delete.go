package cartstorage

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(cartmodel.Cart{}.TableName()).Where(conditions)
	db = db.Delete(&cartmodel.Cart{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
