package entity

import "github.com/dstgo/kratosx/types"

type Env struct {
	Token       string  `json:"token" gorm:"column:token"`
	Keyword     string  `json:"keyword" gorm:"column:keyword"`
	Name        string  `json:"name" gorm:"column:name"`
	Status      *bool   `json:"status" gorm:"column:status"`
	Description *string `json:"description" gorm:"column:description"`
	types.BaseModel
}
