// /model/database/database.go

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tortorCoin/internal/utils"
	"tortorCoin/pkg/config"
)

var db *gorm.DB

// Init 初始化数据库连接
func Init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		utils.Log.Error("Failed to load config: ", err)
		return
	}

	db, err = gorm.Open(mysql.Open(cfg.Database.DSN), &gorm.Config{})
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
