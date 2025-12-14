package eventstorage

import (
	"context"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(eventmodel.Event{}.TableName()).Where(conditions)
	db = db.Delete(&eventmodel.Event{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
