package user

import (
	"github.com/aayushjoshi2709/mypic/src/utils/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup, handler *Handler) {
	group.POST("", handler.create)
	group.POST("/login", handler.login)

	authenticatedGroup := group.Group("")
	authenticatedGroup.Use(middleware.AuthMiddleware)
	authenticatedGroup.GET("/me", handler.getCurrentUser)
	authenticatedGroup.GET("/:id", handler.get)
	authenticatedGroup.GET("", handler.getAll)
	authenticatedGroup.PUT("/:id", handler.update)
	authenticatedGroup.DELETE("/logout", handler.logout)
	authenticatedGroup.DELETE("/:id", handler.delete)
}
