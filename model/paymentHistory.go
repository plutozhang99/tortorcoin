package model

import (
	"errors"
	"gorm.io/gorm"
	"tortorCoin/internal/utils"
	"tortorCoin/model/database"
)

type PaymentHistory struct {
	gorm.Model
	PaymentID  string `gorm:"unique;not null;primaryKey;autoIncrement;index"`
	ReceiverID string `gorm:"not null;index"`
	SenderID   string `gorm:"not null;index"`
	Amount     int64  `gorm:"not null"`
	CreateAt   int64  `gorm:"autoCreateTime"`
	Note       string
}

const PaymentHistoryTableName = "payment_histories"

func (*PaymentHistory) TableName() string {
	return PaymentHistoryTableName
}

// CreatePaymentHistory creates a new payment history, note is not required.
func CreatePaymentHistory(paymentID string, receiverID string, senderID string, amount int64, note string) (*PaymentHistory, error) {
	paymentHistory := &PaymentHistory{
		PaymentID:  paymentID,
		ReceiverID: receiverID,
		SenderID:   senderID,
		Amount:     amount,
		Note:       note,
	}

	if err := database.GetDB().Create(paymentHistory).Error; err != nil {
		utils.Log.Error("Failed to create payment history: ", err)
		return nil, utils.ErrInsertFailed
	}

	return paymentHistory, nil
}

// GetPaymentHistoryByUserID returns a list of payment history by userID.
func GetPaymentHistoryByUserID(userID string) ([]*PaymentHistory, error) {
	var paymentHistories []*PaymentHistory
	if err := database.GetDB().Where("receiver_id = ? or sender_id = ?", userID, userID).Find(&paymentHistories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Log.Error("Payment History not found by User Id: ", userID)
			return nil, utils.ErrPaymentHistoryNotFound
		}
		utils.Log.Error("Failed to get payment history by userID: ", err)
		return nil, utils.ErrGetPaymentFailed
	}

	return paymentHistories, nil
}
