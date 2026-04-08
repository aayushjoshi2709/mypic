package src

import (
	"github.com/aayushjoshi2709/mypic/src/comment"
	"github.com/aayushjoshi2709/mypic/src/image"
	"github.com/aayushjoshi2709/mypic/src/presign"
	"github.com/aayushjoshi2709/mypic/src/user"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(engine *gin.Engine) {
	apiGroup := engine.Group("/api")

	userHandler := GetHandler[user.Handler]("user_handler")
	user.Routes(apiGroup.Group("/v1/user"), userHandler)

	imageHandler := GetHandler[image.Handler]("image_handler")
	image.Routes(apiGroup.Group("/v1/image"), imageHandler)

	commentHandler := GetHandler[comment.Handler]("comment_handler")
	comment.Routes(apiGroup.Group("/v1/comment"), commentHandler)

	presignHandler := GetHandler[presign.Handler]("presign_handler")
	presign.Routes(apiGroup.Group("/v1/presign"), presignHandler)

	engine.Static("/assets", "./public/assets")

	engine.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	engine.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

}
