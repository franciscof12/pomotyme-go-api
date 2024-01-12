package main

import (
	"log"

	"github.com/franciscof12/pomotyme-go-api/v1/initializers"
	"github.com/franciscof12/pomotyme-go-api/v1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initializers.ConnectDatabase()
	// import routes from routes package
	routes.InitRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s\n", err)
	}
}
