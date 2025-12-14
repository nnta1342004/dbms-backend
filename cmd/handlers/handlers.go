package handlers

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	"hareta/middleware"
	blogtaggin "hareta/modules/blog-tag/transport/gin"
	bloggin "hareta/modules/blog/transport/gin"
	cartgin "hareta/modules/cart/transport/gin"
	eventgin "hareta/modules/event/transport/gin"
	eventitemgin "hareta/modules/event_item/transport/gin"
	groupitemgin "hareta/modules/group_item/transport/gin"
	imagegin "hareta/modules/image/transport/gin"
	itemgin "hareta/modules/item/transport/gin"
	itemimagegin "hareta/modules/item_image/transport/gin"
	itemordergin "hareta/modules/item_order/transport/gin"
	ordergin "hareta/modules/order/transport/gin"
	passwordrecoverygin "hareta/modules/password_recovery/transport/gin"
	userstorage "hareta/modules/user/storage"
	usergin "hareta/modules/user/transport/gin"
	userlikeitemgin "hareta/modules/user_like_item/transport/gin"
	"net/http"
)

func MainRoute(r *gin.Engine, sc goservice.ServiceContext) {
	r.MaxMultipartMemory = 15 << 20
	r.Use(middleware.AllowCORS(), middleware.Recover(sc), middleware.RateLimit(sc))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", usergin.Register(sc))

	r.POST("/login", usergin.Login(sc))
	r.POST("/send-verification-code", usergin.SendLink(sc))

	r.GET("/check-verification-code/:id", usergin.CheckLink(sc))

	db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)

	store := userstorage.NewSQLStore(db)

	passwordRecovery := r.Group("/password-recovery")
	{
		passwordRecovery.POST("/", passwordrecoverygin.Create(sc))
		passwordRecovery.POST("/recovery", passwordrecoverygin.RecoverPassword(sc))
		passwordRecovery.GET("/", passwordrecoverygin.Find(sc))
	}

	auth := r.Group("/auth", middleware.RequiredAuth(sc, store))
	{
		auth.GET("/", usergin.GetMe(sc))
		auth.POST("/change-password", usergin.ChangePassword(sc))
		auth.POST("/avatar", usergin.UploadAvatar(sc))
		auth.PUT("/", usergin.Update(sc))
	}

	authItem := auth.Group("/item", middleware.RequiredAdminAuth(sc))
	{
		authItem.PUT("/avatar", itemgin.UpdateAvt(sc))
		authItem.POST("/", itemgin.Create(sc))
		authItem.PUT("/", itemgin.Update(sc))
		authItem.DELETE("/", itemgin.Delete(sc))
		authItem.PUT("/default", itemgin.MakeDefault(sc))
	}

	item := r.Group("/item")
	{
		item.GET("/:id", itemgin.Find(sc))
		item.GET("/", itemgin.List(sc))
		item.GET("/filter", itemgin.ListType(sc))
		item.GET("/group", itemgin.ListItemInGroup(sc))
	}

	itemImage := r.Group("/item-image")
	{
		itemImage.POST("/", middleware.RequiredAuth(sc, store), middleware.RequiredAdminAuth(sc), itemimagegin.AddImages(sc))
		itemImage.GET("/", itemimagegin.List(sc))
		itemImage.DELETE("/", middleware.RequiredAuth(sc, store), middleware.RequiredAdminAuth(sc), itemimagegin.Delete(sc))
		itemImage.PUT("/", middleware.RequiredAuth(sc, store), middleware.RequiredAdminAuth(sc), itemimagegin.Update(sc))
	}

	image := auth.Group("/image", middleware.RequiredAdminAuth(sc))
	{
		image.POST("/", imagegin.UploadByFile(sc))
		image.GET("/", imagegin.List(sc))
		image.DELETE("/", imagegin.Delete(sc))
	}

	cart := auth.Group("/cart")
	{
		cart.POST("/", cartgin.Create(sc))
		cart.DELETE("/", cartgin.Delete(sc))
		cart.GET("/", cartgin.List(sc))
		cart.PUT("/", cartgin.Update(sc))
	}

	userLikeItem := auth.Group("/user-like-item")
	{
		userLikeItem.POST("/", userlikeitemgin.Create(sc))
		userLikeItem.DELETE("/", userlikeitemgin.Delete(sc))
		userLikeItem.GET("/", userlikeitemgin.ListLikedItem(sc))
	}

	order := auth.Group("/order")
	{
		order.POST("/", ordergin.Create(sc))
		order.GET("/", ordergin.List(sc))
		order.PUT("/", middleware.RequiredAdminAuth(sc), ordergin.Update(sc))
		order.GET("/admin", middleware.RequiredAdminAuth(sc), ordergin.ListAdmin(sc))
		order.GET("/admin/:id", middleware.RequiredAdminAuth(sc), ordergin.FindOrderAdmin(sc))
		order.GET("/:id", ordergin.Find(sc))
	}
	r.POST("/order", ordergin.CreateWithoutLogin(sc))
	r.GET("/order/:id", ordergin.FindWithoutLogin(sc))

	itemOrder := auth.Group("/item-order")
	{
		itemOrder.GET("/", itemordergin.List(sc))
	}
	r.GET("/item-order", itemordergin.ListWithoutLogin(sc))

	groupItem := auth.Group("/group-item", middleware.RequiredAdminAuth(sc))
	{
		groupItem.GET("/", groupitemgin.List(sc))
		groupItem.POST("/", groupitemgin.Create(sc))
		groupItem.DELETE("/", groupitemgin.Delete(sc))
		groupItem.PUT("/", groupitemgin.Update(sc))
		groupItem.PUT("/price/:group-id", groupitemgin.UpdatePrice(sc))
	}

	event := r.Group("/event")
	{
		event.GET("/", eventgin.List(sc))
		event.GET("/:id", eventgin.Find(sc))

		authEvent := event.Group("/", middleware.RequiredAuth(sc, store))
		{
			authEvent.POST("/", middleware.RequiredAdminAuth(sc), eventgin.Create(sc))
			authEvent.PUT("/", middleware.RequiredAdminAuth(sc), eventgin.Update(sc))
			authEvent.DELETE("/", middleware.RequiredAdminAuth(sc), eventgin.Delete(sc))
			authEvent.GET("/admin", middleware.RequiredAdminAuth(sc), eventgin.ListAdmin(sc))
		}
	}

	eventItem := r.Group("/event-item")
	{
		eventItem.POST("/", middleware.RequiredAdminAuth(sc), eventitemgin.Create(sc))
		eventItem.DELETE("/", middleware.RequiredAdminAuth(sc), eventitemgin.Delete(sc))
	}

	blog := r.Group("/blog")
	{
		blog.POST(
			"/",
			middleware.RequiredAuth(sc, store),
			middleware.RequiredAdminAuth(sc),
			bloggin.Create(sc),
		)
		blog.GET("/:id", bloggin.Find(sc))
		blog.GET("/", bloggin.List(sc))
		blog.PUT(
			"/",
			middleware.RequiredAuth(sc, store),
			middleware.RequiredAdminAuth(sc),
			bloggin.Update(sc),
		)
		blog.DELETE(
			"/",
			middleware.RequiredAuth(sc, store),
			middleware.RequiredAdminAuth(sc),
			bloggin.Delete(sc),
		)

		tag := blog.Group("/tag")
		{
			tag.POST(
				"/",
				middleware.RequiredAuth(sc, store),
				middleware.RequiredAdminAuth(sc),
				blogtaggin.Create(sc),
			)
			tag.DELETE(
				"/",
				middleware.RequiredAuth(sc, store),
				middleware.RequiredAdminAuth(sc),
				blogtaggin.Delete(sc),
			)
			tag.GET("/", blogtaggin.List(sc))
		}
	}
}
