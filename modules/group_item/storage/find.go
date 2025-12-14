package groupitemstorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*groupitemmodel.GroupItem, error) {
	db := s.db.Table(groupitemmodel.GroupItem{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data groupitemmodel.GroupItem
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
