package main

import (
	"tortorCoin/internal/auth"
	"tortorCoin/model/database"
)

func main() {
	// 初始化数据库连接
	database.Init()
	// TODO: 将error handler放入Auth包中
	err := auth.InitConfig()
	if err != nil {
		return
	}
}
