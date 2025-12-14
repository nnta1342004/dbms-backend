package userstorage

import (
	"context"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
)

func (s *sqlStore) Create(ctx context.Context, user *usermodel.User) error {
	db := s.db
	if err := db.Create(user).Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
