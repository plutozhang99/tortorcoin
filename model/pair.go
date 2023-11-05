package model

import (
	"gorm.io/gorm"
)

type Pair struct {
	gorm.Model
	MatchID     string `gorm:"unique;not null"`
	FirstUser   string `gorm:"not null"`
	SecondUser  string `gorm:"not null"`
	Status      string `gorm:"default:'pending'"`
	CreateAt    int64  `gorm:"autoCreateTime"`
	UpdateAt    int64  `gorm:"autoUpdateTime"`
	IsMissmatch bool   `gorm:"default:false"`
}
