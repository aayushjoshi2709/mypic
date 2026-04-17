package presign

import (
	"fmt"
	"io"
	"log/slog"
	"time"

	"net/http"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/aayushjoshi2709/mypic/src/utils/redis"
	"github.com/gin-gonic/gin"
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
	presignedObj, err := h.repo.PutObject(
		ctx.Request.Context(),
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

// @Summary Get an image by public URL
// @Description Get an image by its public URL
// @Tags Presign
// @Accept json
// @Produce json
// @Param id path string true "Public URL ID"
// @Success 200 {file} file
// @Failure 400 {object} common.ErrorResponseDto
// @Router /api/v1/presign/{id} [get]
func (h *Handler) getImageByPublicUrl(ctx *gin.Context) {
	id := ctx.Param("id")

	val, err := redis.Init().GetAndDelete(ctx.Request.Context(), id)
	if err != nil {
		slog.Error("Error fetching original image key from Redis", "error", err)
		ctx.JSON(http.StatusNotFound, common.ErrorResponseDto{Error: "Image not found"})
		return
	}

	originalImageKey := val

	obj, err := h.repo.GetObjectStream(ctx.Request.Context(), originalImageKey)

	if err != nil {
		slog.Error("Error fetching object from S3", "error", err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "An error occurred while fetching the image"})
		return
	}

	ctx.Header("Content-Type", "image/jpeg")
	ctx.Header("Content-Length", fmt.Sprintf("%d", *obj.ContentLength))
	ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", originalImageKey))
	// 1 day cache control for the image as it will be available in s3 for 1 hour and we want to avoid hitting s3 for every request
	ctx.Header("Cache-Control", "public, max-age=86400")

	bodyCh := make(chan []byte, 1)
    errCh := make(chan error, 1)

    go func() {
        defer obj.Body.Close()
        b, err := io.ReadAll(obj.Body)
        if err != nil {
            errCh <- err
            return
        }
        bodyCh <- b
    }()

    select {
    case err := <-errCh:
        slog.Error("Error reading object body", "error", err)
        ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "An error occurred while streaming the image"})
        return
    case b := <-bodyCh:
        ctx.Data(http.StatusOK, "image/jpeg", b)
    }
}