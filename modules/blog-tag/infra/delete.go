package blogtaginfra

import (
	"context"
	"hareta/appCommon"
	blogtagmodel "hareta/modules/blog-tag/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(blogtagmodel.BlogTag{}.TableName()).Where(conditions)
	db = db.Delete(&blogtagmodel.BlogTag{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
