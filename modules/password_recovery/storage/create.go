package passwordrecoverystorage

import (
	"context"
	"fmt"
	"hareta/appCommon"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
)

func (s *sqlStore) Create(ctx context.Context, user *passwordrecoverymodel.PasswordRecovery) error {
	db := s.db.Begin()
	fmt.Println(user)
	if err := db.Create(user).Error; err != nil {
		db.Rollback()
		return appCommon.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		return appCommon.ErrInternal(err)
	}
	return nil
}
