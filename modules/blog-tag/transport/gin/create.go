package blogtaggin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	blogtaginfra "hareta/modules/blog-tag/infra"
	blogtagusecase "hareta/modules/blog-tag/usecase"
	blogstorage "hareta/modules/blog/infra"
	"net/http"
)

func Create(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data blogtagusecase.BlogTagCreateDTO
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := blogtaginfra.NewSQLStore(db)
		blogStore := blogstorage.NewSQLStore(db)

		biz := blogtagusecase.NewBlogTagCreateUseCase(store, blogStore)
		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
