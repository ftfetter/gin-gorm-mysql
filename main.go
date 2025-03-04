package main

import (
	"gin-gorm-mysql/business"
	"gin-gorm-mysql/config"
	"gin-gorm-mysql/controllers"
	"gin-gorm-mysql/database"
	"gin-gorm-mysql/routes"
	"os"

	_ "gin-gorm-mysql/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	repository := database.NewUserRepository(db.InitDB())
	service := business.NewUserService(repository)
	controller := controllers.NewUserController(service)
	router := routes.NewRouter(controller)

	router.SetupRouter(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(":8080")
}
