package userlikeitemstorage

import (
	"context"
	"hareta/appCommon"
	userlikeitemmodel "hareta/modules/user_like_item/model"
)

func (s *sqlStore) Create(ctx context.Context, data *userlikeitemmodel.UserLikeItem) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
