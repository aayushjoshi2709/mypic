package src

import (
	"github.com/aayushjoshi2709/mypic/src/comment"
	"github.com/aayushjoshi2709/mypic/src/image"
	"github.com/aayushjoshi2709/mypic/src/user"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(engine *gin.Engine)  {
	apiGroup := engine.Group("/api")
	user.Routes(apiGroup.Group("/v1/user"))
	image.Routes(apiGroup.Group("/v1/image"))
	comment.Routes(apiGroup.Group("/v1/comment"))
}
