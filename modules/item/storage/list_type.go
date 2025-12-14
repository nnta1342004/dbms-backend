package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) ListType(ctx context.Context, conditions map[string]interface{}, param string) ([]string, error) {
	db := s.db.Table(itemmodel.Item{}.TableName()).Where(conditions).Where("status != ?", itemmodel.StatusDeleted)
	var result []string
	if err := db.Distinct(param).Pluck(param, &result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
