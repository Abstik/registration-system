package utils

import (
	"registration_system/dao/mysql"
	"registration_system/models"
)

func InitSqlTable() {
	mysql.DB.AutoMigrate(&models.User{})
	mysql.DB.AutoMigrate(&models.Interview{})
}
