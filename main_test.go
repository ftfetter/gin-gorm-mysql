package main

import (
	"gin-gorm-mysql/business"
	"gin-gorm-mysql/config"
	"gin-gorm-mysql/controllers"
	"gin-gorm-mysql/database"
	"gin-gorm-mysql/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func TestMain(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new router
	r := gin.Default()

	// Initialize the database configuration
	db := config.DBConfig{
		Host:     DB_HOST,
		Port:     DB_PORT,
		User:     DB_USERNAME,
		DBName:   DB_NAME,
		Password: DB_PASSWORD,
	}

	// Initialize the repository, service, and controller
	repository := database.NewUserRepository(db.InitDB())
	service := business.NewUserService(repository)
	controller := controllers.NewUserController(service)
	router := routes.NewRouter(controller)

	// Setup the router
	router.SetupRouter(r)

	// Add the Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Create a test server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make a request to the Swagger endpoint
	resp, err := http.Get(ts.URL + "/swagger/index.html")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
