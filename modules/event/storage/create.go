package eventstorage

import (
	"context"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
)

func (s *sqlStore) Create(ctx context.Context, data *eventmodel.Event) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
