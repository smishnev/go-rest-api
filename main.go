package main

import (
	"github.com/gin-gonic/gin"
	"rest-api.com/db"
	"rest-api.com/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
