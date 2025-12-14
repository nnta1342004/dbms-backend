package itemimagebiz

import (
	"context"
	"hareta/appCommon"
	itemimagemodel "hareta/modules/item_image/model"
)

type listStore interface {
	List(ctx context.Context, paging *appCommon.Paging, conditions map[string]interface{}, moreInfo ...string) ([]itemimagemodel.ItemImage, error)
}
type listBiz struct {
	store listStore
}

func NewListBiz(store listStore) *listBiz {
	return &listBiz{store: store}
}
func (biz *listBiz) List(ctx context.Context, data *itemimagemodel.ItemList) ([]itemimagemodel.ItemImage, error) {
	data.Paging.Fulfill()
	id, err := appCommon.FromBase58(data.Id)
	if err != nil {
		return nil, appCommon.ErrInvalidRequest(err)
	}
	res, err := biz.store.List(ctx, &data.Paging, map[string]interface{}{"item_id": id.GetLocalID()}, "Image")

	for i := range res {
		res[i].Mask(false)
		if res[i].Image != nil {
			res[i].Image.Mask(false)
		}
	}
	return res, nil
}
