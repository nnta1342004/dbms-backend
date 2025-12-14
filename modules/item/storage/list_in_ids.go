package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) ListInIds(ctx context.Context, ids []int64, conditions map[string]interface{}, moreInfo ...string) ([]itemmodel.Item, error) {
	db := s.db.Table(itemmodel.Item{}.TableName()).Where("status != ?", itemmodel.StatusDeleted).Where(conditions)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var result []itemmodel.Item

	if err := db.Where("id IN ?", ids).Find(&result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
