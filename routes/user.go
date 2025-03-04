package routes

import (
	"gin-gorm-mysql/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	users controllers.UserController
}

func NewRouter(users controllers.UserController) *Router {
	return &Router{users: users}
}

func (r *Router) SetupRouter(e *gin.Engine) {
	v1 := e.Group("/api/v1")
	{
		users := v1.Group("users")
		{
			users.GET("", r.users.FetchUsers)
			users.POST("", r.users.CreateUser)
			users.GET(":id", r.users.FetchUserById)
			users.PATCH(":id", r.users.UpdateUser)
			users.DELETE(":id", r.users.DeleteUser)
		}
	}
}
