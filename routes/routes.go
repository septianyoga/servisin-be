package routes

import (
	"github.com/gin-gonic/gin"
	"myapp/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}

	return r
}
