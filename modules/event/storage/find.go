package eventstorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*eventmodel.Event, error) {
	db := s.db.Table(eventmodel.Event{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data eventmodel.Event
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
