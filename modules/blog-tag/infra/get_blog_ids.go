package blogtaginfra

import (
	"context"
	"hareta/appCommon"
	blogtagmodel "hareta/modules/blog-tag/model"
)

func (s *sqlStore) GetBlogIds(ctx context.Context, tag string) ([]int64, error) {
	db := *s.db.WithContext(ctx)

	var result []int64
	if err := db.Table(blogtagmodel.BlogTag{}.TableName()).Select("blog_id").Where("tag = ?", tag).Pluck(
		"blog_id",
		&result,
	).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
