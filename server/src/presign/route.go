package presign

import (
	"github.com/aayushjoshi2709/mypic/src/utils/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup, handler *Handler) {
	group.POST("", middleware.AuthMiddleware, handler.getUrl)
	group.GET("/:id", handler.getImageByPublicUrl)
}
