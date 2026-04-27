package group

import (
	"github.com/aayushjoshi2709/mypic/src/utils/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup, handler *Handler) {
	group.Use(middleware.AuthMiddleware)
	group.POST("/", handler.add)
	group.GET("/", handler.getAll)
	group.GET("/:id", handler.get)
	group.DELETE("/:id", handler.delete)
	group.POST("/:id/image", handler.addImage)
	group.POST("/:id/user", handler.addUser)
	group.GET("/:id/image", handler.getImages)
	group.GET("/:id/user", handler.getUsers)
	group.DELETE("/:id/image/:imageId", handler.deleteImage)
	group.DELETE("/:id/user/:userId", handler.deleteUser)
}
