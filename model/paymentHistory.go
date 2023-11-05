package model

import (
	"gorm.io/gorm"
)

type PaymentHistory struct {
	gorm.Model
	PaymentID  string `gorm:"unique;not null"`
	ReceiverID string `gorm:"not null"`
	SenderID   string `gorm:"not null"`
	Amount     int64  `gorm:"not null"`
	CreateAt   int64  `gorm:"autoCreateTime"`
	Note       string
}
