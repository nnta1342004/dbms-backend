package blogusecase

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

type BlogCreateDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Avatar  string `json:"avatar"`
	Overall string `json:"overall"`
}

type IBlogCreateRepo interface {
	Create(ctx context.Context, data *blogmodel.Blog) error
}
type IBlogUsecase interface {
	Create(ctx context.Context, data *BlogCreateDTO) (*string, error)
}

type blogCreateUsecase struct {
	repo   IBlogCreateRepo
	logger logger.Logger
}

func NewBlogCreateUsecase(repo IBlogCreateRepo) IBlogUsecase {
	return &blogCreateUsecase{
		repo:   repo,
		logger: logger.GetCurrent().GetLogger("BlogCreateUseCase"),
	}
}

func (biz *blogCreateUsecase) Create(ctx context.Context, data *BlogCreateDTO) (*string, error) {
	createData := &blogmodel.Blog{
		Title:   data.Title,
		Content: data.Content,
		Avatar:  data.Avatar,
		Overall: data.Overall,
	}

	if err := biz.repo.Create(ctx, createData); err != nil {
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotCreateEntity(blogmodel.EntityName, err)
	}

	return &createData.FakeId, nil
}
