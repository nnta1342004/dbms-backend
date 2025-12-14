package appCommon

import (
	"time"
)

type SQLModelNew struct {
	Id        int64      `json:"-" gorm:"column:id;primaryKey"`
	FakeId    string     `json:"id" gorm:"type:varchar(255);column:fake_id;unique;not null;"`
	Status    int        `json:"status" gorm:"column:status;default:0;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;autoUpdateTime"`
}

const ShardId = 7022004

type SQLModel struct {
	Id        int64      `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:0;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;autoUpdateTime"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, ShardId)
	m.FakeId = &uid
}
