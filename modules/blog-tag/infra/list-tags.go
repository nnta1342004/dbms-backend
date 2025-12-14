package blogtaginfra

import (
	"context"
	"hareta/appCommon"
	blogtagmodel "hareta/modules/blog-tag/model"
)

func (s *sqlStore) ListTags(ctx context.Context) ([]string, error) {
	db := *s.db.WithContext(ctx)

	var result []string
	if err := db.Table(blogtagmodel.BlogTag{}.TableName()).Select("DISTINCT tag").Pluck(
		"tag",
		&result,
	).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}

	return result, nil
}
