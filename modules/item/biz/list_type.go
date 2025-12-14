package itembiz

import (
	"context"
	"errors"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	itemmodel "hareta/modules/item/model"
	"strings"
)

type listTypeStore interface {
	ListType(ctx context.Context, conditions map[string]interface{}, param string) ([]string, error)
}

type listTypeBiz struct {
	store  listTypeStore
	logger logger.Logger
}

func NewListTypeBiz(store listTypeStore) *listTypeBiz {
	return &listTypeBiz{store: store, logger: logger.GetCurrent().GetLogger("ItemListTypeBiz")}
}

func (biz *listTypeBiz) List(ctx context.Context, filter *itemmodel.ItemTypeList) ([]string, error) {
	blocks := strings.Split(filter.Query, "&")
	conditions := make(map[string]interface{})
	for i := range blocks {
		res := strings.Split(blocks[i], ":")
		if len(res) != 2 {
			continue
		}
		conditions[res[0]] = res[1]
	}
	if filter.Field != "category" && filter.Field != "type" && filter.Field != "collection" {
		return nil, appCommon.ErrInvalidRequest(errors.New("cannot list type of this field"))
	}
	res, err := biz.store.ListType(ctx, conditions, filter.Field)
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []string{}, appCommon.ErrCannotListEntity("Type", err)
	}
	return res, nil
}
