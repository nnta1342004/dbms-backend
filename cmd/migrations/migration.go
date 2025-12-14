package main

import (
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/storage/sdkgorm"
	"gorm.io/gorm"
	"hareta/appCommon"
	blogtagmodel "hareta/modules/blog-tag/model"
	blogmodel "hareta/modules/blog/model"
	eventitemmodel "hareta/modules/event_item/model"
)

func main() {
	sc := goservice.New(
		goservice.WithName("hareta"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("gorm", appCommon.PluginGorm)),
	)
	if err := sc.Init(); err != nil {
		panic(err)
	}

	db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
	sc.Logger("migrations").Info("Migrating database")
	if err := db.AutoMigrate(
		//&itemmodel.Item{},
		//&ordermodel.Order{},
		//&itemimagemodel.ItemImage{},
		//&itemordermodel.ItemOrder{},
		//&cartmodel.Cart{},
		//&usermodel.User{},
		//&passwordrecoverymodel.PasswordRecovery{},
		//&userlikeitemmodel.UserLikeItem{},
		//&groupitemmodel.GroupItem{},
		//&eventmodel.Event{},
		&eventitemmodel.EventItem{},
		&blogmodel.Blog{},
		&blogtagmodel.BlogTag{},
	); err != nil {
		panic(err)
	}
}
