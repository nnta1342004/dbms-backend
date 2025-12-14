package blogtagmodel

const EntityName = "BlogTag"

type BlogTag struct {
	BlogId int64  `json:"-" gorm:"column:blog_id;primaryKey;OnDelete:CASCADE"`
	Tag    string `json:"tag" gorm:"type:varchar(255);column:tag;primaryKey"`
}

func (BlogTag) TableName() string {
	return "blog_tag"
}
