package itemstorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) FindDataById(ctx context.Context, id int64, moreInfo ...string) (*itemmodel.Item, error) {
	db := s.db.Table(itemmodel.Item{}.TableName()).
		Where("id = ?", id).
		Where("status != ?", itemmodel.StatusDeleted)

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
