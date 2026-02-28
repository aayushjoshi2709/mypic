package routes

import (
	"github.com/aayushjoshi2709/mypic/routes/comment"
	"github.com/aayushjoshi2709/mypic/routes/image"
	"github.com/aayushjoshi2709/mypic/routes/user"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(engine *gin.Engine)  {
	apiGroup := engine.Group("/api")
	user.Routes(apiGroup.Group("/v1/user"))
	image.Routes(apiGroup.Group("/v1/image"))
	comment.Routes(apiGroup.Group("/v1/comment"))
}
