package model

import (
	"github.com/LynnWonder/gin_prac/pkg/db"
	"time"
)

type BaseModel struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;type:bigint" json:"id"`                                                     // id
	CreatedAt time.Time `gorm:"column:created_at;not null;type:TIMESTAMP;index:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`            // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
}

func init() {
	db.Models = append(db.Models, &Person{})
}
