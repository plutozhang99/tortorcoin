package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tortorCoin/internal/utils"
)

var db *gorm.DB

// Init 初始化数据库连接
func Init() {
	dsn := "root:1234@tcp(localhost:3306)/TorTorCoin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log.Error("Failed to connect to database: ", err)
	} else {
		utils.Log.Info("Database connection established")
	}
}

// GetDB 返回数据库连接实例
func GetDB() *gorm.DB {
	return db
}
