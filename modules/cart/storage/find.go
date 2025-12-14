package cartstorage

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*cartmodel.Cart, error) {
	db := s.db.Table(cartmodel.Cart{}.TableName()).Where(condition)
	for i := range moreInfo {
		fmt.Println(moreInfo[i])
		db = db.Preload(moreInfo[i])
	}
	var data cartmodel.Cart
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
