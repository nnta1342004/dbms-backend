package blogusecase

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	blogtagmodel "hareta/modules/blog-tag/model"
	blogmodel "hareta/modules/blog/model"
)

type BlogListDTO struct {
	Title *string `json:"title" form:"title"`
	Tag   *string `json:"tag" form:"tag"`
}

type blogListRepo interface {
	List(
		ctx context.Context,
		paging *appCommon.Paging,
		filter *BlogListDTO,
		ids []int64,
		moreInfo ...string,
	) ([]blogmodel.SimpleBlog, error)
}

type blogTagListRepo interface {
	GetBlogIds(ctx context.Context, tag string) ([]int64, error)
}

type blogListUseCase interface {
	List(ctx context.Context, paging *appCommon.Paging, filter *BlogListDTO) ([]blogmodel.SimpleBlog, error)
}

type blogListUsecase struct {
	repo        blogListRepo
	blogTagRepo blogTagListRepo
	logger      logger.Logger
}

func NewBlogListUseCase(repo blogListRepo, blogTagRepo blogTagListRepo) blogListUseCase {
	return &blogListUsecase{
		repo:        repo,
		blogTagRepo: blogTagRepo,
		logger:      logger.GetCurrent().GetLogger("BlogListUseCase"),
	}
}

func (biz *blogListUsecase) List(ctx context.Context, paging *appCommon.Paging, filter *BlogListDTO) ([]blogmodel.SimpleBlog, error) {
	paging.Fulfill()
	var ids []int64

	if filter.Tag != nil {
		res, err := biz.blogTagRepo.GetBlogIds(ctx, *filter.Tag)
		if err != nil {
			biz.logger.WithSrc().Errorln(err)
			return []blogmodel.SimpleBlog{}, appCommon.ErrCannotListEntity(blogtagmodel.EntityName, err)
		}
		ids = res
	} else {
		ids = nil
	}

	result, err := biz.repo.List(ctx, paging, filter, ids, "Tags")
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return []blogmodel.SimpleBlog{}, appCommon.ErrCannotListEntity(blogmodel.EntityName, err)
	}
	if result != nil && len(result) > 0 {
		paging.NextCursor = result[len(result)-1].FakeId
	}
	return result, nil
}
