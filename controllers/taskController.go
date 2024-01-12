package controllers

import (
	"net/http"

	"github.com/franciscof12/pomotyme-go-api/v1/initializers"
	models "github.com/franciscof12/pomotyme-go-api/v1/models/mysql"
	"github.com/franciscof12/pomotyme-go-api/v1/schemas"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks, err := models.GetAllTasks(initializers.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := models.GetTaskByID(initializers.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func CreateTask(c *gin.Context) {
	var input schemas.TaskSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{Category: input.Category, Description: input.Description, UserID: input.UserID, Completed: input.Completed}
	err := models.CreateTask(initializers.DB, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var input schemas.TaskSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := models.GetTaskByID(initializers.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task.Category = input.Category
	task.Description = input.Description
	task.UserID = input.UserID
	task.Completed = input.Completed
	err = models.UpdateTask(initializers.DB, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := models.GetTaskByID(initializers.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteTask(initializers.DB, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
