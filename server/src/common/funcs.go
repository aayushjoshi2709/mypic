package common

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPageAndLimit(ctx *gin.Context) (int64, int64, error) {
	pageInt, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		slog.Error("Error converting page variable", "error", err)
		ctx.JSON(http.StatusBadRequest, ErrorResponseDto{Error: "The value provided for page is not valid"})
		return 0, 0, err
	}

	limitInt, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		slog.Error("Error converting limit variable", "error", err)
		ctx.JSON(http.StatusBadRequest, ErrorResponseDto{Error: "The value provided for limit is not valid"})
		return 0, 0, err
	}
	return int64(pageInt), int64(limitInt), nil
}
