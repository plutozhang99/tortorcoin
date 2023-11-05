package model

import (
	"errors"
	"gorm.io/gorm"
	"tortorCoin/internal/utils"
	"tortorCoin/model/database"
)

type User struct {
	gorm.Model
	UserID     uint   `gorm:"unique;not null;autoIncrement;primaryKey;index"`
	Password   string `gorm:"not null"`
	UserName   string `gorm:"unique;not null;index"`
	Account    string `gorm:"unique;not null;index"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
	UpdatedAt  int64  `gorm:"autoUpdateTime"`
	MatchID    uint   `gorm:"default:''"`
	CoinAmount int64  `gorm:"default:5"`
	IsDeleted  bool   `gorm:"default:false"`
}

// ToMap converts User object to a map representation.
func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"userID":     u.UserID,
		"userName":   u.UserName,
		"account":    u.Account,
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
func CreateUser(userName string, password string, account string) (*User, error) {
	user := &User{
		Password: password,
		UserName: userName,
		Account:  account,
	}

	if err := database.GetDB().Create(user).Error; err != nil {
		utils.Log.Error("Failed to create user: ", err)
		return nil, err
	}

	return user, nil
}

// GetUserByField returns a user by a given field and value.
func GetUserByField(field string, value interface{}) (*User, error) {
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
	return GetUserByField("user_name", userName)
}

// GetUserByUserID returns a user by userID.
func GetUserByUserID(userID uint) (*User, error) {
	return GetUserByField("id", userID)
}

// GetUserByAccount returns a user by account.
func GetUserByAccount(account string) (*User, error) {
	return GetUserByField("account", account)
}

// UpdateUser updates a user.
func UpdateUser(userID uint, user *User) error {
	if err := database.GetDB().Model(&User{}).Where("id = ?", userID).Updates(user).Error; err != nil {
		utils.Log.Error("Failed to update user: ", err)
		return err
	}
	return nil
}

// CheckUsernameExist checks if a user exists by userName. Returns true if user exists.
func CheckUsernameExist(userName string) bool {
	// Call GetUserByUserName to check if user exists.
	user, _ := GetUserByUserName(userName)
	if user != nil {
		return true
	}
	return false
}

// CheckAccountExist checks if a user exists by account. Returns true if user exists.
func CheckAccountExist(account string) bool {
	// Call GetUserByAccount to check if user exists.
	user, _ := GetUserByAccount(account)
	if user != nil {
		return true
	}
	return false
}
