package userlikeitemstorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	userlikeitemmodel "hareta/modules/user_like_item/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*userlikeitemmodel.UserLikeItem, error) {
	db := s.db.Table(userlikeitemmodel.UserLikeItem{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data userlikeitemmodel.UserLikeItem
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
