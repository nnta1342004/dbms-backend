package eventitemstorage

import (
	"context"
	"hareta/appCommon"
	eventitemmodel "hareta/modules/event_item/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(eventitemmodel.EventItem{}.TableName()).Where(conditions)
	db = db.Delete(&eventitemmodel.EventItem{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
