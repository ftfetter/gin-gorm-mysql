package database

import (
	"fmt"
	"gin-gorm-mysql/config"
	"gin-gorm-mysql/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "my_db"
const DB_HOST = "127.0.0.1"
const DB_PORT = 9910

func InitDb() *gorm.DB {
	db := ConnectDB()
	db.AutoMigrate(&models.User{})
	return db
}

func ConnectDB() *gorm.DB {
	conf := config.DBConfig{
		Host:     DB_HOST,
		Port:     DB_PORT,
		User:     DB_USERNAME,
		DBName:   DB_NAME,
		Password: DB_PASSWORD,
	}

	db, err := gorm.Open(mysql.Open(conf.MysqlUrl()), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database: error=%v\n", err)
	}

	return db
}
