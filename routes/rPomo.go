package routes

import (
	"github.com/franciscof12/pomotyme-go-api/v1/controllers"
	"github.com/gin-gonic/gin"
)

func PomoRoutes(r *gin.Engine) {
	pomo := r.Group("/pomos")
	{
		pomo.GET("/", controllers.GetPomodoros)
		pomo.GET("/:id", controllers.GetPomodoroByID)
		pomo.POST("/", controllers.CreatePomodoro)
		pomo.PUT("/:id", controllers.UpdatePomodoro)
		pomo.DELETE("/:id", controllers.DeletePomodoro)
	}
}
