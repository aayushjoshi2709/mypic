package user

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup) {
	group.GET("/", get)
	group.GET("/:id", getAll)
	group.POST("/", create)
	group.PUT("/:id", update)
	group.DELETE("/:id", delete)
	group.POST("/login", login)
	group.POST("/logout", logout)
}
