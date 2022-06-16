package model
import "time"

type Person struct {
	Id                   int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name                 string     `json:"name" gorm:"uniqueIndex;size:100;not null"`
	Description          *string    `json:"description" gorm:"size:200"`
	CreatedTime          time.Time  `json:"createdTime" gorm:"autoCreateTime"`
	UpdatedTime          time.Time  `json:"updatedTime" gorm:"autoUpdateTime"`
}