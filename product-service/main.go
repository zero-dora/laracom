package main

import "github.com/zero-dora/laracom/user-service/db"

func main() {
	// 创建数据库连接，程序退出时断开连接
	database, err := db.CreateConnection()
	defer database.Close()
}
