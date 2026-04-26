package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/aayushjoshi2709/mypic/docs"
	"github.com/aayushjoshi2709/mypic/src"
	"github.com/aayushjoshi2709/mypic/src/utils/db"
	"github.com/aayushjoshi2709/mypic/src/utils/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found. The server will be using values provided in the os environment variables.")
	}
	db.Init()
}

// @title MyPic API
// @version 1.0
// @description API for MyPic application
// @host localhost:3000
// @basePath /
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// cors

	router.Use(func(c *gin.Context) {
		if os.Getenv("ENV") != "production" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
			c.Next()
		}
	})

	src.SetUpRepositories(&gin.Context{})
	src.SetUpHandlers()
	router.Use(middleware.RequestIdMiddleware)
	src.SetUpRoutes(router)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("PORT environment variable not set, defaulting to 8080")
		port = 8080
	}

	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
