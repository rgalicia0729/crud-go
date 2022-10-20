package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rgalicia0729/crud-go/src/infrastructure/db"
	"github.com/rgalicia0729/crud-go/src/infrastructure/router"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Connect to the database
	db.PostgresConnect(os.Getenv("POSTGRES_URI"))

	gin.SetMode(gin.ReleaseMode)
	server := gin.New()

	// Middlewares are added
	server.Use(gin.Logger())

	// Routes are added
	router.InitRoutes(server)

	server.Run(":8080")
}
