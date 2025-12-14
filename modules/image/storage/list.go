package imagestorage

import (
	"context"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
)

func (s *sqlStore) List(
	ctx context.Context,
	paging *appCommon.Paging,
	filter *imagemodel.ImageFilter,
	conditions map[string]interface{},
	moreInfo ...string,
) ([]imagemodel.Image, error) {
	db := s.db.Table(imagemodel.Image{}.TableName()).Where(conditions)
	db = db.Order("id DESC")
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		if uid, err := appCommon.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	result := make([]imagemodel.Image, 0)

	db = db.Where("created_at BETWEEN ? AND ?", filter.TimeFrom, filter.TimeTo)

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
