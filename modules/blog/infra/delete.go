package blogstorage

import (
	"context"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(blogmodel.Blog{}.TableName()).Where(conditions)
	db = db.Delete(&blogmodel.Blog{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
