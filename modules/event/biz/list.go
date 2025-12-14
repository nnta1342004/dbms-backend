package eventbiz

import (
	"context"
	"fmt"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	eventmodel "hareta/modules/event/model"
	usermodel "hareta/modules/user/model"
)

type listStore interface {
	List(
		ctx context.Context,
		paging *appCommon.Paging,
		conditions map[string]interface{},
		isAdmin bool,
		moreInfo ...string,
	) ([]eventmodel.Event, error)
}
type listBiz struct {
	store  listStore
	logger logger.Logger
}

func NewListBiz(store listStore) *listBiz {
	return &listBiz{
		store:  store,
		logger: logger.GetCurrent().GetLogger("EventListBiz"),
	}
}

func (biz *listBiz) List(ctx context.Context, paging *appCommon.Paging, user *usermodel.User) ([]eventmodel.Event, error) {
	if paging == nil {
		paging = &appCommon.Paging{
			Page:  1,
			Limit: 50,
		}
	}

	fmt.Println(user)

	paging.Fulfill()
	res, err := biz.store.List(
		ctx,
		paging,
		nil,
		user != nil && user.Role == "admin",
		"Items",
		"Items.Item",
		"Items.Item.Avatar",
	)
	if err != nil {
		biz.logger.Errorln(err)
		return nil, appCommon.ErrCannotListEntity(eventmodel.EntityName, err)
	}
	if len(res) == 0 {
		return []eventmodel.Event{}, nil
	}
	for i := range res {
		res[i].Mask(false)
	}
	for i := range res {
		for j := range res[i].Items {
			res[i].Items[j].Item.Mask(false)
		}
	}
	paging.NextCursor = res[len(res)-1].FakeId.String()
	return res, nil
}
