package groupitemstorage

import (
	"context"
	"hareta/appCommon"
	groupitemmodel "hareta/modules/group_item/model"
)

func (s *sqlStore) List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]groupitemmodel.GroupItem, error) {
	db := s.db.Table(groupitemmodel.GroupItem{}.TableName()).Where(conditions)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		if uid, err := appCommon.FromBase58(v); err == nil {
			db = db.Where("id > ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	result := make([]groupitemmodel.GroupItem, 0)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	if err := db.
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
