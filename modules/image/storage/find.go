package imagestorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	imagemodel "hareta/modules/image/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*imagemodel.Image, error) {
	db := s.db.Table(imagemodel.Image{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data imagemodel.Image
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
