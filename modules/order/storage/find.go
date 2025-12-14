package orderstorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	ordermodel "hareta/modules/order/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*ordermodel.Order, error) {
	db := s.db.Table(ordermodel.Order{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data ordermodel.Order
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
