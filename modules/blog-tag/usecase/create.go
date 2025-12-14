package blogtagusecase

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	blogtagmodel "hareta/modules/blog-tag/model"
	blogmodel "hareta/modules/blog/model"
)

type IBlogTagCreateRepo interface {
	Create(ctx context.Context, data []blogtagmodel.BlogTag) error
}

type IBlogCreateRepo interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*blogmodel.Blog, error)
}

type BlogTagCreateDTO struct {
	BlogId string   `json:"blog_id" binding:"required"`
	Tag    []string `json:"tag" binding:"required"`
}
type IBlogTagUseCase interface {
	Create(ctx context.Context, dto *BlogTagCreateDTO) error
}

type blogTagCreateUseCase struct {
	blogTagRepo IBlogTagCreateRepo
	blogRepo    IBlogCreateRepo
	logger      logger.Logger
}

func NewBlogTagCreateUseCase(
	blogTagRepo IBlogTagCreateRepo,
	blogRepo IBlogCreateRepo,
) IBlogTagUseCase {
	return &blogTagCreateUseCase{
		blogTagRepo: blogTagRepo,
		blogRepo:    blogRepo,
		logger:      logger.GetCurrent().GetLogger("BlogTagCreateUseCase"),
	}
}

func (biz *blogTagCreateUseCase) Create(ctx context.Context, dto *BlogTagCreateDTO) error {
	blog, err := biz.blogRepo.FindDataWithCondition(ctx, map[string]interface{}{"fake_id": dto.BlogId})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(blogmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(blogmodel.EntityName, err)
	}

	var data []blogtagmodel.BlogTag
	for i := range dto.Tag {
		data = append(data, blogtagmodel.BlogTag{
			BlogId: blog.Id,
			Tag:    dto.Tag[i],
		})
	}

	if err := biz.blogTagRepo.Create(ctx, data); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotCreateEntity(blogtagmodel.EntityName, err)
	}
	return nil
}
