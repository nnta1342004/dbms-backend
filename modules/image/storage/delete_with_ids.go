package imagestorage

import (
	"context"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
)

func (s *sqlStore) DeleteWithIds(ctx context.Context, ids []int64) error {
	db := s.db.Table(imagemodel.Image{}.TableName()).Where("id IN ?", ids)
	db = db.Delete(&imagemodel.Image{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
