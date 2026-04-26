package group

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup, handler *Handler){
	group.POST("/", handler.add)
	group.GET("/", handler.get)
	group.GET("/:id", handler.get)
	group.DELETE("/:id", handler.delete)
	group.POST("/:id/image", handler.addImage)
	group.POST("/:id/user", handler.addUser)
}