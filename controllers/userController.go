package controllers

import (
	"net/http"

	"github.com/franciscof12/pomotyme-go-api/v1/app"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := c.MustGet("databaseConn").(app.Repository)
	users, err := db.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("databaseConn").(app.Repository)
	user, err := db.GetUserByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("databaseConn").(app.Repository)
	user, err := db.DeleteUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("databaseConn").(app.Repository)
	user, err := db.UpdateUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
	db := c.MustGet("databaseConn").(app.Repository)
	user, err := db.CreateUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
