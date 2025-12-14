package passwordrecoverystorage

import (
	"context"
	"hareta/appCommon"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
)

func (s *sqlStore) UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(passwordrecoverymodel.PasswordRecovery{}.TableName()).Where(conditions).Updates(data)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
