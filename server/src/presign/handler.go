package presign

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/aayushjoshi2709/mypic/src/utils/store"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	typeBucketMap map[string]string
}

func (h *Handler) New() {
	h.typeBucketMap = map[string]string{
		"images": "mypic-images",
		"profilePic": "mypic-videos",
	}
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
	s3Presigner, err := store.New(ctx.Request.Context())
	if err != nil {
		slog.Error("Error creating S3 presigner", "error", err)
		ctx.JSON(500, common.ErrorResponseDto{Error: "An error occurred while creating the S3 presigner"})
		return
	}

	var presignedObjectRequest = &PresignedObjectRequest{}
	if err := ctx.ShouldBindBodyWithJSON(presignedObjectRequest); err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "Provided data is not valid"})
		return
	}

	key := fmt.Sprintf("image-%d-%s", time.Now().Unix(), presignedObjectRequest.OriginalName)
	bucket := h.typeBucketMap[presignedObjectRequest.Type]

	url, err := s3Presigner.PutObject(
		ctx.Request.Context(),
		bucket,
		key,
		15,
	)
	if err != nil {
		slog.Error("Error generating presigned URL", "error", err)
		ctx.JSON(500, common.ErrorResponseDto{Error: "An error occurred while generating the presigned URL"})
		return
	}

	ctx.JSON(200, PresignedObjectResponse{
		URL: url,
		BucketName: bucket,
		Key: key,
	})
}
