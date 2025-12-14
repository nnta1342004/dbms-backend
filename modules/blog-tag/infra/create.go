package blogtaginfra

import (
	"context"
	"hareta/appCommon"
	blogtagmodel "hareta/modules/blog-tag/model"
)

func (s *sqlStore) Create(ctx context.Context, data []blogtagmodel.BlogTag) error {
	db := *s.db
	if err := db.Create(&data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
