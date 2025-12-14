package orderstorage

import (
	"context"
	"hareta/appCommon"
	ordermodel "hareta/modules/order/model"
)

func (s *sqlStore) UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(ordermodel.Order{}.TableName()).Where(conditions).Updates(data)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
