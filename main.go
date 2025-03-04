package main

import (
	"gin-gorm-mysql/business"
	"gin-gorm-mysql/config"
	"gin-gorm-mysql/controllers"
	"gin-gorm-mysql/database"
	"gin-gorm-mysql/routes"

	_ "gin-gorm-mysql/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TODO: Get these values from environment variables
const (
	DB_USERNAME = "root"
	DB_PASSWORD = "root"
	DB_NAME     = "my_db"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = 9910
)

// @title           Gin + GORM + MySQL API
// @version         1.0
// @description     Application to practice GoLang using Gin, GORM and MySQL.

// @license.name	Apache 2.0
// @license.url	    http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	r := gin.Default()

	db := config.DBConfig{
		Host:     DB_HOST,
		Port:     DB_PORT,
		User:     DB_USERNAME,
		DBName:   DB_NAME,
		Password: DB_PASSWORD,
	}

	repository := database.NewUserRepository(db.InitDB())
	service := business.NewUserService(repository)
	controller := controllers.NewUserController(service)
	router := routes.NewRouter(controller)

	router.SetupRouter(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(":8080")
}
