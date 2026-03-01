package comment

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup, handler *Handler) {
	group.GET("/", handler.get)
	group.GET("/:id", handler.getAll)
	group.POST("/", handler.create)
	group.PUT("/:id", handler.update)
	group.DELETE("/:id", handler.delete)
}