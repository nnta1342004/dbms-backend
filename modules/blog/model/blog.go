package blogmodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hareta/appCommon"
	blogtagmodel "hareta/modules/blog-tag/model"
)

const EntityName = "Blog"

type Blog struct {
	appCommon.SQLModelNew `json:",inline"`
	Content               string                 `json:"content" gorm:"type:text;column:content"`
	Avatar                string                 `json:"avatar" gorm:"type:varchar(255);column:avatar"`
	Title                 string                 `json:"title" gorm:"type:varchar(255);column:title"`
	Overall               string                 `json:"overall" gorm:"type:text;column:overall"`
	Tags                  []blogtagmodel.BlogTag `json:"tags" gorm:"foreignKey:BlogId;references:Id;OnDelete:CASCADE"`
}

type SimpleBlog struct {
	appCommon.SQLModelNew `json:",inline"`
	Avatar                string                 `json:"avatar" gorm:"column:avatar"`
	Title                 string                 `json:"title" gorm:"column:title"`
	Overall               string                 `json:"overall" gorm:"column:overall"`
	Tags                  []blogtagmodel.BlogTag `json:"tags" gorm:"foreignKey:BlogId;references:Id;OnDelete:CASCADE"`
}

func (u *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.FakeId = id.String()
	return
}

func (Blog) TableName() string {
	return "blog"
}
func (SimpleBlog) TableName() string { return "blog" }
