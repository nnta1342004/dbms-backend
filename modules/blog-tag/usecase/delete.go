package blogtagusecase

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

type IBlogTagDeleteRepo interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
}

type IBlogDeleteRepo interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*blogmodel.Blog, error)
}

type blogTagDeleteUseCase struct {
	blogTagRepo IBlogTagDeleteRepo
	blogRepo    IBlogDeleteRepo
	logger      logger.Logger
}

func NewBlogTagDeleteUseCase(
	blogTagRepo IBlogTagDeleteRepo,
	blogRepo IBlogDeleteRepo,
) *blogTagDeleteUseCase {
	return &blogTagDeleteUseCase{
		blogTagRepo: blogTagRepo,
		blogRepo:    blogRepo,
		logger:      logger.GetCurrent().GetLogger("BlogTagDeleteUseCase"),
	}
}

type BlogTagDeleteDTO struct {
	BlogID string `json:"blog_id" binding:"required"`
	Tag    string `json:"tag" binding:"required"`
}

func (biz *blogTagDeleteUseCase) DeleteBlogTag(ctx context.Context, data *BlogTagDeleteDTO) error {
	blog, err := biz.blogRepo.FindDataWithCondition(ctx, map[string]interface{}{"fake_id": data.BlogID})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(blogmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(blogmodel.EntityName, err)
	}

	if err := biz.blogTagRepo.DeleteWithCondition(ctx, map[string]interface{}{
		"blog_id": blog.Id,
		"tag":     data.Tag,
	}); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(blogmodel.EntityName, err)
	}

	return nil
}
