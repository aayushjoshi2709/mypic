package image

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/gin-gonic/gin"

	"net/http"
)

type Handler struct {
	repos         map[string]any
	cloudFrontUrl string
}

func (h *Handler) New(repos map[string]any) {
	h.repos = repos
	h.cloudFrontUrl = os.Getenv("AWS_CLOUD_FRONT_URL")
	if h.cloudFrontUrl == "" {
		slog.Error("Cloudfront url not found")
		panic("Cloudfront url not found")
	}
}

// @Summary Get an image by ID
// @Description Get an image by its unique ID
// @Tags Images
// @Accept json
// @Produce json
// @Param id path string true "Image ID"
// @Success 200 {object} GetImageResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/image/{id} [get]
func (h *Handler) get(ctx *gin.Context) {
	id := ctx.Param("id")
	image, err := h.repos["imageRepository"].(*Repository).GetById(ctx, id)
	if err != nil {
		slog.Error(fmt.Sprintf("Error getting image with id: %s", id), "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while getting the image"})
		return
	}
	var getImageResponse GetImageResponse
	getImageResponse.Set(ctx, h.cloudFrontUrl, image)
	ctx.JSON(http.StatusOK, getImageResponse)
}

// @Summary Get all images
// @Description Get all images in the database
// @Tags Images
// @Accept json
// @Produce json
// @Param page query number false "Page number" default(1)
// @Param limit query number false "Number of images per page" default(10)
// @Success 200 {object} []GetImageResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/image [get]
func (h *Handler) getAll(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		slog.Error("Error converting page variable", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "The value provided for page is not valid"})
		return
	}


	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		slog.Error("Error converting limit variable", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "The value provided for limit is not valid"})
		return
	}

	images, err := h.repos["imageRepository"].(*Repository).GetAll(ctx, pageInt, limitInt)

	GetImageResponseArr := []GetImageResponse{}
	for _, image := range images {
		var getImageResponse GetImageResponse
		getImageResponse.Set(ctx, h.cloudFrontUrl, &image)
		GetImageResponseArr = append(GetImageResponseArr, getImageResponse)
	}
	ctx.JSON(http.StatusOK, GetImageResponseArr)
}

// @Summary Create a new image
// @Description Create a new image with the provided details
// @Tags Images
// @Accept json
// @Produce json
// @Param image body CreateImageRequest true "Image details"
// @Success 201 {object} GetImageResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/image [post]
func (h *Handler) create(ctx *gin.Context) {
	createImageRequest := &CreateImageRequest{}
	err := ctx.ShouldBindBodyWithJSON(createImageRequest)

	if err != nil {
		slog.Error("Error creating image entity", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Invalid key name provided"})
		return
	}

	err = h.repos["imageRepository"].(*Repository).Add(ctx, createImageRequest.Key, createImageRequest.OriginalName)

	if err != nil {
		slog.Error("Error creating image", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while creating the image"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Image created successfully"})
}

// @UpdateImage godoc
// @Summary Update an existing image
// @Description Update an existing image with the provided details
// @Tags Images
// @Accept json
// @Produce json
// @Param id path string true "Image ID"
// @Param image body UpdateImageRequest true "Updated image details"
// @Success 200 {object} GetImageResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/image/{id} [put]
func (h *Handler) update(ctx *gin.Context) {
	UpdateImageRequest := &UpdateImageRequest{}
	if err := ctx.ShouldBindBodyWithJSON(UpdateImageRequest); err != nil {
		slog.Error("Error updating image", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Provided data is not valid"})
		return
	}

	image, err := h.repos["imageRepository"].(*Repository).Update(ctx, ctx.Param("id"), UpdateImageRequest.Key, UpdateImageRequest.OriginalName)

	if err != nil {
		slog.Error("Error updating image", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while updating the image"})
		return
	}

	var getImageResponse GetImageResponse
	getImageResponse.Set(ctx, h.cloudFrontUrl, image)
	ctx.JSON(200, getImageResponse)
}

// @Summary Delete an existing image
// @Description Delete an existing image by ID
// @Tags Images
// @Accept json
// @Produce json
// @Param id path string true "Image ID"
// @Success 204 "No Content"
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/image/{id} [delete]
func (h *Handler) delete(ctx *gin.Context) {
	err := h.repos["imageRepository"].(*Repository).Delete(ctx, ctx.Param("id"))

	if err != nil {
		slog.Error("Error deleting image", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while deleting the image"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
