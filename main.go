package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aayushjoshi2709/mypic/src"
	"github.com/aayushjoshi2709/mypic/src/utils/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error locating the .env file")
	}
	db.Init()
}

func main() {
	router := gin.Default()
	src.SetUpRoutes(router)
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("PORT environment variable not set, defaulting to 8080")
		port = "8080"
	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
