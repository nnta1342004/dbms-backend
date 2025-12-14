package blogusecase

import (
	"context"
	"github.com/leductoan3082004/go-sdk/logger"
	"hareta/appCommon"
	blogmodel "hareta/modules/blog/model"
)

type BlogFindDTO struct {
	Id string `json:"id" binding:"required"`
}

type IBlogFindRepo interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*blogmodel.Blog, error)
}

type IBlogFindUsecase interface {
	Find(ctx context.Context, data *BlogFindDTO) (*blogmodel.Blog, error)
}

type blogFindUsecase struct {
	repo   IBlogFindRepo
	logger logger.Logger
}

func NewBlogFindUsecase(repo IBlogFindRepo) IBlogFindUsecase {
	return &blogFindUsecase{
		repo:   repo,
		logger: logger.GetCurrent().GetLogger("BlogFindUseCase"),
	}
}

func (biz *blogFindUsecase) Find(ctx context.Context, data *BlogFindDTO) (*blogmodel.Blog, error) {
	condition := map[string]interface{}{
		"fake_id": data.Id,
	}
	res, err := biz.repo.FindDataWithCondition(ctx, condition, "Tags")
	if err != nil {
		if err == appCommon.RecordNotFound {
			return nil, appCommon.ErrEntityNotFound(blogmodel.EntityName, err)
		}
		biz.logger.WithSrc().Errorln(err)
		return nil, appCommon.ErrCannotGetEntity(blogmodel.EntityName, err)
	}
	return res, nil
}
