package image

import (
	"github.com/aayushjoshi2709/mypic/src/utils/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup, handler *Handler) {
	group.Use(middleware.AuthMiddleware)
	group.GET("/:id", handler.get)
	group.GET("", handler.getAll)
	group.POST("", handler.create)
	group.PUT("/:id", handler.update)
	group.DELETE("/:id", handler.delete)
}
