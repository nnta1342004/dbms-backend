package bloggin

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"gorm.io/gorm"
	"hareta/appCommon"
	blogstorage "hareta/modules/blog/infra"
	blogusecase "hareta/modules/blog/usecase"
	"net/http"
)

func Find(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		data := blogusecase.BlogFindDTO{Id: id}
		db := sc.MustGet(appCommon.PluginGorm).(*gorm.DB)
		store := blogstorage.NewSQLStore(db)
		biz := blogusecase.NewBlogFindUsecase(store)
		res, err := biz.Find(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
