package cartstorage

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
)

func (s *sqlStore) List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]cartmodel.Cart, error) {
	db := s.db.Table(cartmodel.Cart{}.TableName()).Where(conditions).Order("id desc")
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		if uid, err := appCommon.FromBase58(v); err == nil {
			db = db.Where("id > ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	result := make([]cartmodel.Cart, 0)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	if err := db.
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
