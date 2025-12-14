package itemstorage

import (
	"context"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
)

func (s *sqlStore) List(
	ctx context.Context,
	filter *itemmodel.ItemList,
	paging *appCommon.Paging,
	moreInfo ...string,
) ([]itemmodel.SimpleItem, error) {
	db := s.db.Table(itemmodel.SimpleItem{}.TableName()).Where("status != ?", itemmodel.StatusDeleted).
		Where("`default` = ?", 1).Order("id DESC")
	if filter != nil {
		typeList := filter.TypeFilter
		if typeList.Category != nil {
			db = db.Where("category = ?", *typeList.Category)
		}
		if typeList.Collection != nil {
			db = db.Where("collection = ?", *typeList.Collection)
		}
		if typeList.Type != nil {
			db = db.Where("type = ?", *typeList.Type)
		}
		if typeList.ProductLine != nil {
			db = db.Where("product_line = ?", *typeList.ProductLine)
		}
		if typeList.Tag != nil {
			db = db.Where("tag = ?", *typeList.Tag)
		}

		priceList := filter.PriceFilter
		if priceList.LowerPrice != nil {
			db = db.Where("price >= ?", *priceList.LowerPrice)
		}
		if priceList.UpperPrice != nil {
			db = db.Where("price <= ?", *priceList.UpperPrice)
		}
		if priceList.Desc != nil && *priceList.Desc {
			db = db.Order("price DESC")
		}

		searchList := filter.SearchFilter
		if searchList.Name != nil {
			db = db.Where("name LIKE ?", "%"+*searchList.Name+"%")
		}
	}
	if paging != nil {
		if err := db.Count(&paging.Total).Error; err != nil {
			return []itemmodel.SimpleItem{}, appCommon.ErrDB(err)
		}

		db = db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit)
	}

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var result []itemmodel.SimpleItem
	if err := db.Find(&result).Error; err != nil {
		return []itemmodel.SimpleItem{}, appCommon.ErrDB(err)
	}

	if result == nil {
		return []itemmodel.SimpleItem{}, nil
	}
	return result, nil
}
