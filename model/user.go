package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID     string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	UserName   string `gorm:"unique;not null"`
	NickName   string
	CreatedAt  int64 `gorm:"autoCreateTime"`
	UpdatedAt  int64 `gorm:"autoUpdateTime"`
	MatchID    string
	CoinAmount int64 `gorm:"default:5"`
	IsDeleted  bool  `gorm:"default:false"`
}
