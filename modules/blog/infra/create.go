package blogstorage

import (
	"context"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

func (s *sqlStore) Create(ctx context.Context, data *blogmodel.Blog) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
