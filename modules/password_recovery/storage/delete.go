package passwordrecoverystorage

import (
	"context"
	"hareta/appCommon"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
)

func (s *sqlStore) DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(passwordrecoverymodel.PasswordRecovery{}.TableName()).Where(conditions)
	db = db.Delete(&passwordrecoverymodel.PasswordRecovery{})
	if db.Error != nil {
		return appCommon.ErrDB(db.Error)
	}
	return nil
}
