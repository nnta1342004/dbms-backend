package imagestorage

import (
	"context"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(imagemodel.Image{}.TableName()).Where(conditions)
	db = db.Delete(&imagemodel.Image{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
