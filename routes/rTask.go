package routes

import (
	"github.com/franciscof12/pomotyme-go-api/v1/controllers"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	task := r.Group("/tasks")
	{
		task.GET("/", controllers.GetTasks)
		task.GET("/:id", controllers.GetTaskByID)
		task.POST("/", controllers.CreateTask)
		task.PUT("/:id", controllers.UpdateTask)
	}
}
