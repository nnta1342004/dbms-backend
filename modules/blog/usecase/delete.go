package blogusecase

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

type BlogDeleteDTO struct {
	Id string `json:"id" binding:"required"`
}

type IBlogDeleteRepo interface {
	DeleteWithCondition(ctx context.Context, conditions map[string]interface{}) error
}

type IBlogDeleteUseCase interface {
	Delete(ctx context.Context, dto *BlogDeleteDTO) error
}

type blogDeleteUseCase struct {
	repo   IBlogDeleteRepo
	logger logger.Logger
}

func NewBlogDeleteUseCase(repo IBlogDeleteRepo) IBlogDeleteUseCase {
	return &blogDeleteUseCase{
		repo:   repo,
		logger: logger.GetCurrent().GetLogger("BlogDeleteUseCase"),
	}
}

func (biz *blogDeleteUseCase) Delete(ctx context.Context, dto *BlogDeleteDTO) error {
	err := biz.repo.DeleteWithCondition(ctx, map[string]interface{}{"fake_id": dto.Id})
	if err != nil {
		biz.logger.WithSrc().Errorln(err)
		return appCommon.ErrCannotDeleteEntity(blogmodel.EntityName, err)
	}
	return nil
}
