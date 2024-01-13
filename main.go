package main

import (
	"log"

	"github.com/franciscof12/pomotyme-go-api/v1/app"

	"github.com/franciscof12/pomotyme-go-api/v1/initializers"
	"github.com/franciscof12/pomotyme-go-api/v1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repo := initializers.ConnectMysqlDatabase()
	r.Use(ApiMiddleware(repo))
	routes.UserRoutes(r)
	routes.TaskRoutes(r)
	routes.PomoRoutes(r)
	if err := r.Run(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s\n", err)
	}
}

func ApiMiddleware(db app.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}
