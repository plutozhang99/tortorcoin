package model

import (
	"errors"
	"gorm.io/gorm"
	"tortorCoin/internal/utils"
	"tortorCoin/model/database"
)

type User struct {
	gorm.Model
	UserID     string `gorm:"unique;not null;autoIncrement;primaryKey"`
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
		// 处理错误
		utils.Log.Error("Failed to hash password: ", err)
		return nil, err
	}

	user := &User{
		Password: hashedPassword,
		UserName: userName,
		NickName: nickName,
	}

	err = database.GetDB().Create(user).Error
	if err != nil {
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
	var user User // 用于存储找到的用户

	err := database.GetDB().Table(UserTableName).Where("user_name = ?", userName).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有找到记录
			utils.Log.Info("User not found: ", userName)
			return nil, gorm.ErrRecordNotFound // 或者您可以返回自定义的错误，比如 ErrUserNotFound
		}
		// 查询过程中发生了其他错误
		utils.Log.Error("Failed to retrieve user: ", err)
		return nil, err
	}

	return &user, nil // 返回找到的用户
}

// GetUserByUserID returns a user by userID.
func GetUserByUserID(userID string) (*User, error) {
	var user User

	err := database.GetDB().Table(UserTableName).Where("user_id=?", userID).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有找到记录
			utils.Log.Info("User not found: ", userID)
			return nil, gorm.ErrRecordNotFound // 或者您可以返回自定义的错误，比如 ErrUserNotFound
		}
		// 查询过程中发生了其他错误
		utils.Log.Error("Failed to retrieve user: ", err)
		return nil, err
	}
	return &user, nil // 返回找到的用户
}

// UpdateUser updates a user.
func UpdateUser(username string, user *User) error {
	return database.GetDB().Table(UserTableName).Where("user_name = ?", username).Updates(user).Error
}
