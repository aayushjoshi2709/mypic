package comment

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup, handler *Handler) {
	group.GET("/:id", handler.get)
	group.GET("", handler.getAll)
	group.POST("", handler.create)
	group.PUT("/:id", handler.update)
	group.DELETE("/:id", handler.delete)
}
