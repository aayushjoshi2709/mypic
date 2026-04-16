package presign

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	typeBucketMap map[string]string;
	repo *Repository
}

func (h *Handler) New(repository *Repository) {
	h.repo = repository;
}


// @Summary Get presigned URL for uploading an object to S3
// @Description  Get presigned URL for uploading an object to S3
// @Tags Presign
// @Accept json
// @Produce json
// @Param presignedObjectRequest body PresignedObjectRequest true "Presigned Object Request"
// @Success 200 {object} PresignedObjectResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Router /api/v1/presign [post]
func (h *Handler) getUrl(ctx *gin.Context) {
	var presignedObjectRequest = &PresignedObjectRequest{}
	if err := ctx.ShouldBindBodyWithJSON(presignedObjectRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Provided data is not valid"})
		return
	}

	key := fmt.Sprintf("image-%d-%s", time.Now().Unix(), presignedObjectRequest.OriginalName)
	url, err := h.repo.PutObject(
		ctx.Request.Context(),
		key,
		15,
	)
	if err != nil {
		slog.Error("Error generating presigned URL", "error", err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "An error occurred while generating the presigned URL"})
		return
	}

	ctx.JSON(http.StatusOK, PresignedObjectResponse{
		URL: url,
	})
}
