package model

import (
	"errors"
	"gorm.io/gorm"
	"tortorCoin/internal/utils"
	"tortorCoin/model/database"
)

type Pair struct {
	gorm.Model
	MatchID        string `gorm:"unique;not null;index;primaryKey;autoIncrement"`
	FirstUserName  string `gorm:"not null; index"`
	SecondUserName string `gorm:"not null; index"`
	Status         string `gorm:"default:'pending'"`
	CreateAt       int64  `gorm:"autoCreateTime"`
	UpdateAt       int64  `gorm:"autoUpdateTime"`
	IsUnpaired     bool   `gorm:"default:false"`
}

// ToMap converts Pair object to a map representation.
func (p *Pair) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"matchID":    p.MatchID,
		"firstUser":  p.FirstUserName,
		"secondUser": p.SecondUserName,
		"status":     p.Status,
		"createAt":   p.CreateAt,
		"updateAt":   p.UpdateAt,
		"isUnpaired": p.IsUnpaired,
	}
}

const PairTableName = "pairs"

func (*Pair) TableName() string {
	return PairTableName
}

// CreatePair creates a new pair.
func CreatePair(firstUserName string, secondUserName string) (*Pair, error) {
	pair := &Pair{
		FirstUserName:  firstUserName,
		SecondUserName: secondUserName,
		Status:         string(utils.Pending),
	}

	if err := database.GetDB().Create(pair).Error; err != nil {
		return nil, err
	}

	return pair, nil
}

// GetPairByMatchID returns a pair that is not unpaired by matchID.
func GetPairByMatchID(matchID string) (*Pair, error) {
	pair := &Pair{}
	if err := database.GetDB().Where("match_id = ? and is_unpaired = ? and status = ?", matchID, false, utils.Success).First(pair).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Log.Error("Failed to get pair by matchID: ", err)
			return nil, gorm.ErrRecordNotFound
		} else {
			utils.Log.Error("Failed to get pair by matchID: ", err)
			return nil, err
		}
	}

	return pair, nil
}

// GetPairByUserName returns a pair by userName.
func GetPairByUserName(userName string) (*Pair, error) {
	pair := &Pair{}
	if err := database.GetDB().Where("(first_user_name = ? or second_user_name = ?) and is_unpaired = ?", userName, userName, false).First(pair).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Log.Error("Failed to get pair by userName: ", err)
			return nil, gorm.ErrRecordNotFound
		} else {
			utils.Log.Error("Failed to get pair by userName: ", err)
			return nil, err
		}
	}

	return pair, nil
}

// GetPairStatusByMatchID returns a pair of status by matchID.
func GetPairStatusByMatchID(matchID string) (string, error) {
	pair := &Pair{}
	if err := database.GetDB().Where("match_id = ?", matchID).First(pair).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Log.Error("Failed to get pair status by matchID: ", err)
			return "", gorm.ErrRecordNotFound
		} else {
			utils.Log.Error("Failed to get pair status by matchID: ", err)
			return "", err
		}
	}
	return pair.Status, nil
}

// UpdatePairStatus updates a pair status.
func UpdatePairStatus(matchID string, status string) error {
	if err := database.GetDB().Model(&Pair{}).Where("match_id = ?", matchID).Update("status", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Log.Error("Failed to update pair status: ", err)
			return gorm.ErrRecordNotFound
		}
		utils.Log.Error("Failed to update pair status: ", err)
		return err
	}

	return nil
}

// UpdatePairIsUnpaired updates a pair isUnpaired.
func UpdatePairIsUnpaired(matchID string, isUnpaired bool) error {
	if err := database.GetDB().Model(&Pair{}).Where("match_id = ?", matchID).Update("is_unpaired", isUnpaired).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Log.Error("Failed to update pair isUnpaired: ", err)
			return gorm.ErrRecordNotFound
		}
		utils.Log.Error("Failed to update pair isUnpaired: ", err)
		return err
	}

	return nil
}
