package presign

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup, handler *Handler) {
	group.POST("/getUrl", handler.getUrl)
}
