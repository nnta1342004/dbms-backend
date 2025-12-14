package cartstorage

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
)

func (s *sqlStore) ListInIds(ctx context.Context, ids []int64, moreInfo ...string) ([]cartmodel.Cart, error) {
	db := s.db.Table(cartmodel.Cart{}.TableName())
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var result []cartmodel.Cart

	if err := db.Where("id IN ?", ids).Find(&result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
