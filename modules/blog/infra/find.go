package blogstorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreInfo ...string,
) (*blogmodel.Blog, error) {
	db := s.db.Table(blogmodel.Blog{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data blogmodel.Blog
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
