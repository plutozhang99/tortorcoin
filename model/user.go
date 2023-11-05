package model

import (
	"errors"
	"gorm.io/gorm"
	"tortorCoin/internal/utils"
	"tortorCoin/model/database"
)

type User struct {
	gorm.Model
	UserID     uint   `gorm:"unique;not null;autoIncrement;primaryKey"`
	Password   string `gorm:"not null"`
	UserName   string `gorm:"unique;not null;index"`
	NickName   string
	CreatedAt  int64  `gorm:"autoCreateTime"`
	UpdatedAt  int64  `gorm:"autoUpdateTime"`
	MatchID    string `gorm:"default:'';index"`
	CoinAmount int64  `gorm:"default:5"`
	IsDeleted  bool   `gorm:"default:false"`
}

// ToMap converts User object to a map representation.
func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"userID":     u.UserID,
		"userName":   u.UserName,
		"nickName":   u.NickName,
		"createdAt":  u.CreatedAt,
		"updatedAt":  u.UpdatedAt,
		"matchID":    u.MatchID,
		"coinAmount": u.CoinAmount,
		"isDeleted":  u.IsDeleted,
	}
}

const UserTableName = "users"

// TableName overrides the table name used by User to `users`
func (*User) TableName() string {
	return UserTableName
}

// CreateUser creates a new user.
func CreateUser(password string, userName string, nickName string) (*User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		utils.Log.Error("Failed to hash password: ", err)
		return nil, err
	}

	user := &User{
		Password: hashedPassword,
		UserName: userName,
		NickName: nickName,
	}

	if err := database.GetDB().Create(user).Error; err != nil {
		utils.Log.Error("Failed to create user: ", err)
		return nil, err
	}

	return user, nil
}

// getUserByField returns a user by a given field and value.
func getUserByField(field string, value interface{}) (*User, error) {
	var user User

	if err := database.GetDB().Where(field+" = ?", value).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Log.Info("User not found: ", value)
			return nil, gorm.ErrRecordNotFound
		}
		utils.Log.Error("Failed to retrieve user: ", err)
		return nil, err
	}

	return &user, nil
}

// GetUserByUserName returns a user by userName.
func GetUserByUserName(userName string) (*User, error) {
	return getUserByField("user_name", userName)
}

// GetUserByUserID returns a user by userID.
func GetUserByUserID(userID uint) (*User, error) {
	return getUserByField("id", userID)
}

// UpdateUser updates a user.
func UpdateUser(userID uint, user *User) error {
	if err := database.GetDB().Model(&User{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		utils.Log.Error("Failed to update user: ", err)
		return err
	}
	return nil
}
