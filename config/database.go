package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MYSQL_URL_PATTERN = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func (config *DBConfig) InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.mysqlUrl()), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database: error=%v\n", err)
	}

	return db
}

func (config *DBConfig) mysqlUrl() string {
	return fmt.Sprintf(
		MYSQL_URL_PATTERN,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}
