package eventstorage

import (
	"context"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
)

func (s *sqlStore) UpdateEvent(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(eventmodel.Event{}.TableName()).Where(conditions).Updates(data)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
