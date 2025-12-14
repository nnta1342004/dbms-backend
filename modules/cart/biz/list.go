package cartbiz

import (
	"context"
	"hareta/appCommon"
	cartmodel "hareta/modules/cart/model"
	usermodel "hareta/modules/user/model"
)

type listStore interface {
	List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]cartmodel.Cart, error)
}

type listBiz struct {
	store listStore
}

func NewListBiz(store listStore) *listBiz {
	return &listBiz{store: store}
}
func (biz *listBiz) List(ctx context.Context, user *usermodel.User, paging *appCommon.Paging) ([]cartmodel.Cart, error) {
	if paging == nil {
		paging = &appCommon.Paging{
			Page:  1,
			Limit: 50,
		}
	}
	paging.Fulfill()
	res, err := biz.store.List(ctx, paging, map[string]interface{}{"user_id": user.Id}, "Item", "Item.Avatar")
	if err != nil {
		return nil, appCommon.ErrCannotListEntity(cartmodel.EntityName, err)
	}
	for i := range res {
		res[i].Mask(false)
		res[i].Item.Mask(false)
		if res[i].Item.Avatar != nil {
			res[i].Item.Avatar.Mask(false)
		}
	}
	if len(res) > 0 {
		paging.NextCursor = res[len(res)-1].FakeId.String()
	}

	return res, nil
}
