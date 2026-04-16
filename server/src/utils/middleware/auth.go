package middleware

import (
	"net/http"
	"strings"

	"github.com/aayushjoshi2709/mypic/src/utils/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth format"})
		ctx.Abort()
		return
	}

	tokenString := parts[1]

	claims, err := jwt.Init().ValidateToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		ctx.Abort()
		return
	}

	ctx.Set("username", claims.Username)
	ctx.Set("userId", claims.UserId.Hex())
	ctx.Next()
}