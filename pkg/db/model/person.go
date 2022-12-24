package model

type Person struct {
	BaseModel
	Name        string  `json:"name" gorm:"type:varchar(64);uniqueIndex;not null;comment:name"`
	Age         int     `json:"age" gorm:"type:int;not null;comment:age"`
	Description *string `json:"description" gorm:"type:varchar(255);not null;comment:description"`
}
