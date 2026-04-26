package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/aayushjoshi2709/mypic/src/utils/jwt"
	"github.com/aayushjoshi2709/mypic/src/utils/redis"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	redisVal, err := redis.Init().Get(ctx, "logged_out_"+tokenString)
	if err == nil && redisVal != "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been invalidated"})
		ctx.Abort()
		return
	}

	claims, err := jwt.Init().ValidateToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		ctx.Abort()
		return
	}

	ctx.Set("username", claims.Subject)
	userId := claims.ID
	userIdObject, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		slog.Error("Error converting userId to ObjectID", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occoured"})
	}

	ctx.Set("userId", userIdObject)
	ctx.Set("token", tokenString)
	ctx.Next()
}
