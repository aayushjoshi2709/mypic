package src

import (
	"github.com/aayushjoshi2709/mypic/src/comment"
	"github.com/aayushjoshi2709/mypic/src/group"
	"github.com/aayushjoshi2709/mypic/src/image"
	"github.com/aayushjoshi2709/mypic/src/presign"
	"github.com/aayushjoshi2709/mypic/src/user"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(engine *gin.Engine) {
	apiGroup := engine.Group("/api")

	userHandler := GetHandler[user.Handler]("userHandler")
	user.Routes(apiGroup.Group("/v1/user"), userHandler)

	imageHandler := GetHandler[image.Handler]("imageHandler")
	image.Routes(apiGroup.Group("/v1/image"), imageHandler)

	commentHandler := GetHandler[comment.Handler]("commentHandler")
	comment.Routes(apiGroup.Group("/v1/comment"), commentHandler)

	presignHandler := GetHandler[presign.Handler]("presignHandler")
	presign.Routes(apiGroup.Group("/v1/presign"), presignHandler)

	groupHandler := GetHandler[group.Handler]("groupHandler")
	group.Routes(apiGroup.Group("/v1/group"), groupHandler)

	engine.Static("/assets", "./public/assets")

	engine.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	engine.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

}
