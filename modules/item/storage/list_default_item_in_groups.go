package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) ListDefaultItemInGroups(ctx context.Context, groupIds []int64, conditions map[string]interface{}, moreInfo ...string) ([]itemmodel.SimpleItem, error) {
	db := s.db.Table(itemmodel.Item{}.TableName()).Where("status != ?", itemmodel.StatusDeleted).Where(conditions)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var result []itemmodel.SimpleItem

	if err := db.Where("group_id IN ?", groupIds).Find(&result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
