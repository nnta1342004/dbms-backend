package eventstorage

import (
	"context"
	"fmt"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
	"time"
)

func (s *sqlStore) List(
	ctx context.Context,
	paging *appCommon.Paging,
	conditions map[string]interface{},
	isAdmin bool,
	moreInfo ...string,
) ([]eventmodel.Event, error) {
	db := s.db.Table(eventmodel.Event{}.TableName()).Where(conditions).Order("id DESC")
	if !isAdmin {
		db = db.Where("? BETWEEN date_start AND date_end", time.Now())
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		if uid, err := appCommon.FromBase58(v); err == nil {
			db = db.Where(fmt.Sprintf("%s.id < ?", eventmodel.Event{}.TableName()), uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var result []eventmodel.Event

	if err := db.
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
