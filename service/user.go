package service

import (
	"fmt"
	"math/rand"
	"tortorCoin/internal/utils"
	"tortorCoin/model"
)

// TODO: Add Authentication

// Register 调用/model/user.go中的User结构体，还有createUser方法做用户注册
func Register(username string, password string, account string) (*model.User, error) {
	for {
		randomNumber := rand.Intn(10000)
		uniqueUsername := fmt.Sprintf("%s#%04d", username, randomNumber)
		if !model.CheckUsernameExist(uniqueUsername) {
			utils.Log.Info("Creating user: " + uniqueUsername)
			encryptedPassword, err := utils.HashPassword(password)
			if err != nil {
				utils.Log.Error("Failed to hash password: ", err)
				return nil, err
			}
			user, err := model.CreateUser(uniqueUsername, encryptedPassword, account)
			if err != nil {
				utils.Log.Error("Failed to create user: "+user.UserName+" ", err)
				return nil, err
			}
			return user, nil
		}
	}
}

// Login 调用/model/user.go中的User结构体，还有getUserByUserName方法做用户登录
func Login(account string, password string) (*model.User, error) {
	user, err := model.GetUserByAccount(account)
	if err != nil {
		utils.Log.Error("Failed to get user by userName: ", err)
		return nil, utils.ErrGetUserFailed
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		utils.Log.Info("Wrong password")
		return nil, utils.ErrWrongPassword
	}

	// Generate JWT Token
	// TODO: Generate JWT Token
	token, err := GenerateJWTToken(user)
	if err != nil {
		utils.Log.Error("Failed to generate JWT Token: ", err)
		return nil, err
	}

	return token, nil
}
