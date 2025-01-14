package main

import (
	"leiloa/db"
	"leiloa/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Println("Error getting the .env variables")
	}

	if os.Getenv("PRODUCTION") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	routes.StartRoutes(r)

	db.ConnectDB()

	db.CreateDB()

	r.Run(":8000")
}
