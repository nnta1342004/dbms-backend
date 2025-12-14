package eventitemstorage

import (
	"context"
	"fmt"
	"hareta/appCommon"
	eventitemmodel "hareta/modules/event_item/model"
)

func (s *sqlStore) List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]eventitemmodel.EventItem, error) {
	db := s.db.Table(eventitemmodel.EventItem{}.TableName()).Where(conditions)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		if uid, err := appCommon.FromBase58(v); err == nil {
			db = db.Where(fmt.Sprintf("%s.id > ?", eventitemmodel.EventItem{}.TableName()), uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	result := make([]eventitemmodel.EventItem, 0)
	for i := range moreInfo {
		db = db.Joins(moreInfo[i])
	}
	if err := db.
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
