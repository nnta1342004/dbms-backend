package blogstorage

import (
	"context"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

func (s *sqlStore) Save(ctx context.Context, data *blogmodel.Blog) error {
	db := s.db
	if err := db.Save(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
