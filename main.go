package main

import (
	"gin-gorm-mysql/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	_ = r.Run(":8080")
}
