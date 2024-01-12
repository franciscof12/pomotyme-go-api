package controllers

import (
	"net/http"

	"github.com/franciscof12/pomotyme-go-api/v1/initializers"
	models "github.com/franciscof12/pomotyme-go-api/v1/models/mysql"
	"github.com/franciscof12/pomotyme-go-api/v1/schemas"
	"github.com/gin-gonic/gin"
)

func GetPomodoros(c *gin.Context) {
	pomodoros, err := models.GetAllPomodoros(initializers.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pomodoros})
}

func GetPomodoroByID(c *gin.Context) {
	id := c.Param("id")
	pomodoro, err := models.GetPomodoroByID(initializers.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pomodoro})
}

func CreatePomodoro(c *gin.Context) {
	var input schemas.PomodoroSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pomodoro := models.Pomodoro{UserID: input.UserID, StudyTime: input.StudyTime, BreakTime: input.BreakTime, EndDate: input.EndDate}
	err := models.CreatePomodoro(initializers.DB, &pomodoro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pomodoro})
}

func UpdatePomodoro(c *gin.Context) {
	id := c.Param("id")
	var input schemas.PomodoroSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pomodoro, err := models.GetPomodoroByID(initializers.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pomodoro.UserID = input.UserID
	pomodoro.StudyTime = input.StudyTime
	pomodoro.BreakTime = input.BreakTime
	pomodoro.EndDate = input.EndDate
	err = models.UpdatePomodoro(initializers.DB, &pomodoro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pomodoro})
}

func DeletePomodoro(c *gin.Context) {
	id := c.Param("id")
	pomodoro, err := models.GetPomodoroByID(initializers.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = models.DeletePomodoro(initializers.DB, &pomodoro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
