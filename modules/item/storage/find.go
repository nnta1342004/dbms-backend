package itemstorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*itemmodel.Item, error) {
	db := s.db.Table(itemmodel.Item{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data itemmodel.Item
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
