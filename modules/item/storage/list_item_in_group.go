package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) ListItemInGroup(
	ctx context.Context,
	paging *appCommon.Paging,
	groupId int64,
	moreInfo ...string,
) ([]itemmodel.SimpleItem, error) {
	db := s.db.Table(itemmodel.Item{}.TableName()).
		Where("group_id = ?", groupId).
		Where("status != ?", itemmodel.StatusDeleted)

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
	result := make([]itemmodel.SimpleItem, 0)
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
