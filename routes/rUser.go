package routes

import (
	"github.com/franciscof12/pomotyme-go-api/v1/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.GET("/", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserByID)
		user.POST("/", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
	}
}
