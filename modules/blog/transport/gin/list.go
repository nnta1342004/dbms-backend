package bloggin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	blogtaginfra "hareta/modules/blog-tag/infra"
	blogstorage "hareta/modules/blog/infra"
	blogusecase "hareta/modules/blog/usecase"
	"net/http"
)

func List(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data blogusecase.BlogListDTO
		if err := c.ShouldBind(&data); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}
		var paging appCommon.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(appCommon.ErrInvalidRequest(err))
		}

		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := blogstorage.NewSQLStore(db)
		blogTagRepo := blogtaginfra.NewSQLStore(db)
		biz := blogusecase.NewBlogListUseCase(store, blogTagRepo)
		res, err := biz.List(c.Request.Context(), &paging, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, paging, nil))
	}
}
