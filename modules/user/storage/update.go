package userstorage

import (
	"context"
	"hareta/appCommon"
	usermodel "hareta/modules/user/model"
)

func (s *sqlStore) UpdateWithCondition(ctx context.Context, conditions map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(usermodel.User{}.TableName()).Where(conditions).Updates(data)
	if err := db.Error; err != nil {
		return appCommon.ErrDB(err)
	}
	return nil
}
