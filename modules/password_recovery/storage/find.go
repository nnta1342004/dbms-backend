package passwordrecoverystorage

import (
	"context"
	"gorm.io/gorm"
	"hareta/appCommon"
	passwordrecoverymodel "hareta/modules/password_recovery/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*passwordrecoverymodel.PasswordRecovery, error) {
	db := s.db.Table(passwordrecoverymodel.PasswordRecovery{}.TableName()).Where(condition)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data passwordrecoverymodel.PasswordRecovery
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.RecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
