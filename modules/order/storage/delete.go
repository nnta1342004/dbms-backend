package orderstorage

import (
	"context"
	"hareta/appCommon"
	ordermodel "hareta/modules/order/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(ordermodel.Order{}.TableName()).Where(conditions)
	db = db.Delete(&ordermodel.Order{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
