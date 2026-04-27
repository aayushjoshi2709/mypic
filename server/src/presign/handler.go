package presign

import (
	"log/slog"

	"net/http"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	typeBucketMap map[string]string
	repos         map[string]any
}

func (h *Handler) New(repos map[string]any) {
	h.repos = repos
}

// @Summary Get presigned URL for uploading an object to S3
// @Description  Get presigned URL for uploading an object to S3
// @Tags Presign
// @Accept json
// @Produce json
// @Param presignedObjectRequest body PresignedObjectRequest true "Presigned Object Request"
// @Success 200 {object} PresignedObjectResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/presign [post]
func (h *Handler) getUrl(ctx *gin.Context) {
	var presignedObjectRequest = &PresignedObjectRequest{}
	if err := ctx.ShouldBindBodyWithJSON(presignedObjectRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Provided data is not valid"})
		return
	}

	key := uuid.New().String()
	presignedObj, err := h.repos["presignRepository"].(*Repository).PutObject(
		ctx,
		key,
		15,
	)
	if err != nil {
		slog.Error("Error generating presigned URL", "error", err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "An error occurred while generating the presigned URL"})
		return
	}

	ctx.JSON(http.StatusOK, presignedObj)
}
