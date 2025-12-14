package imagestorage

import (
	"context"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
)

func (s *sqlStore) Create(ctx context.Context, data *imagemodel.Image) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
