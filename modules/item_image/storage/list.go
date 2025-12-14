package itemimagestorage

import (
	"context"
	"hareta/appCommon"
	itemimagemodel "hareta/modules/item_image/model"
)

func (s *sqlStore) List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]itemimagemodel.ItemImage, error) {
	db := s.db.Table(itemimagemodel.ItemImage{}.TableName()).Where(conditions)
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
	result := make([]itemimagemodel.ItemImage, 0)
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
