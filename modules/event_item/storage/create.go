package eventitemstorage

import (
	"context"
	"hareta/appCommon"
	eventitemmodel "hareta/modules/event_item/model"
)

func (s *sqlStore) Create(ctx context.Context, data *eventitemmodel.EventItem) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
