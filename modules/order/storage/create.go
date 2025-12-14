package orderstorage

import (
	"context"
	"hareta/appCommon"
	ordermodel "hareta/modules/order/model"
)

func (s *sqlStore) Create(ctx context.Context, data *ordermodel.Order) error {
	db := s.db
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
