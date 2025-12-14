package blogusecase

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

type BlogUpdateDTO struct {
	Id      string  `json:"id" binding:"required"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Avatar  *string `json:"avatar"`
	Overall *string `json:"overall"`
}

type IBlogUpdateRepo interface {
	Save(ctx context.Context, data *blogmodel.Blog) error
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*blogmodel.Blog, error)
}

type IBlogUpdateUseCase interface {
	Update(ctx context.Context, dto *BlogUpdateDTO) error
}

type blogUpdateUseCase struct {
	repo   IBlogUpdateRepo
	logger logger.Logger
}

func NewBlogUpdateUseCase(repo IBlogUpdateRepo) IBlogUpdateUseCase {
	return &blogUpdateUseCase{
		repo:   repo,
		logger: logger.GetCurrent().GetLogger("BlogUpdateUseCase"),
	}
}

func (biz *blogUpdateUseCase) Update(ctx context.Context, dto *BlogUpdateDTO) error {
	blog, err := biz.repo.FindDataWithCondition(ctx, map[string]interface{}{"fake_id": dto.Id})
	if err != nil {
		if err == appCommon.RecordNotFound {
			return appCommon.ErrEntityNotFound(blogmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotGetEntity(blogmodel.EntityName, err)
	}
	if dto.Title != nil {
		blog.Title = *dto.Title
	}
	if dto.Content != nil {
		blog.Content = *dto.Content
	}
	if dto.Avatar != nil {
		blog.Avatar = *dto.Avatar
	}
	if dto.Overall != nil {
		blog.Overall = *dto.Overall
	}

	if err := biz.repo.Save(ctx, blog); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotUpdateEntity(blogmodel.EntityName, err)
	}
	return nil
}
