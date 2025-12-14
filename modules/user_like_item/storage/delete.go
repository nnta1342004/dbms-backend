package userlikeitemstorage

import (
	"context"
	"hareta/appCommon"
	userlikeitemmodel "hareta/modules/user_like_item/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(userlikeitemmodel.UserLikeItem{}.TableName()).Where(conditions)
	db = db.Delete(&userlikeitemmodel.UserLikeItem{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
