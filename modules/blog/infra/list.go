package blogstorage

import (
	"context"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
	blogusecase "hareta/modules/blog/usecase"
)

func (s *sqlStore) List(
	ctx context.Context,
	paging *appCommon.Paging,
	filter *blogusecase.BlogListDTO,
	ids []int64,
	moreInfo ...string,
) ([]blogmodel.SimpleBlog, error) {
	db := s.db.Table(blogmodel.SimpleBlog{}.TableName()).Order("id DESC")

	if filter != nil {
		if v := filter.Title; v != nil && *v != "" {
			db = db.Where("title LIKE ?", "%"+*v+"%")
		}
	}

	if ids != nil {
		db = db.Where("id IN (?)", ids)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	if v := paging.FakeCursor; v != "" {
		db = db.Where("fake_id < ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	result := make([]blogmodel.SimpleBlog, 0)
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	if err := db.
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}
	return result, nil
}
