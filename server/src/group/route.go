package group

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup, handler *Handler) {
	group.POST("/", handler.add)
	group.GET("/", handler.getAll)
	group.GET("/:id", handler.get)
	group.DELETE("/:id", handler.delete)
	group.POST("/:id/image", handler.addImage)
	group.POST("/:id/user", handler.addUser)
	group.GET("/:id/image", handler.getImages)
	group.GET("/:id/user", handler.getUsers)
	group.DELETE("/:id/image", handler.deleteImage)
	group.DELETE("/:id/user", handler.deleteUser)
}
