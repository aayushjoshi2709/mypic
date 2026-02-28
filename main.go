package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/aayushjoshi2709/mypic/routes"
)


func init(){
	err:= godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error locating the .env file")
	}
}

func main(){
	router := gin.Default()

	routes.SetUpRoutes(router)

	err := router.Run(fmt.Sprintf(": %s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}