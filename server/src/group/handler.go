package group

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/aayushjoshi2709/mypic/src/image"
	"github.com/aayushjoshi2709/mypic/src/user"
	"github.com/gin-gonic/gin"
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

func (h *Handler) add(ctx *gin.Context) {
	createGroupRequest := &CreateGroupRequest{}
	err := ctx.ShouldBindBodyWithJSON(createGroupRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Invalid value for the group name"})
	}

	err = h.repos["groupRepository"].(*Repository).Add(ctx, createGroupRequest.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "Something went wrong"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Group created successfully"})
}

func (h *Handler) get(ctx *gin.Context) {
	id := ctx.Param("id")
	group, err := h.repos["groupRepository"].(*Repository).GetById(ctx, id)
	if err != nil {
		slog.Error(fmt.Sprintf("Error getting group with id: %s", id), "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while getting the image"})
		return
	}
	var getGroupResponse GetGroupResponse
	getGroupResponse.Set(ctx, group)
	ctx.JSON(http.StatusOK, getGroupResponse)
}

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
	groups, err := h.repos["groupRepository"].(*Repository).GetAll(ctx, pageInt, limitInt)
	getGroupResponseArr := []GetGroupResponse{}

	for _, group := range groups {
		var getGroupResponse GetGroupResponse
		getGroupResponse.Set(ctx, &group)
		getGroupResponseArr = append(getGroupResponseArr, getGroupResponse)
	}
	ctx.JSON(http.StatusOK, getGroupResponseArr)

}

func (h *Handler) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.repos["groupRepository"].(*Repository).Delete(ctx, id)
	if err != nil {
		slog.Error("Error deleting group", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while deleting the group"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) addImage(ctx *gin.Context) {
	groupId := ctx.Param("id")
	addImageRequest := &AddImageRequest{}
	err := ctx.ShouldBindBodyWithJSON(addImageRequest)
	if err != nil {
		slog.Error("Error parsing the payload for add image request: ", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Please provide image id in request body"})
		return
	}

	imageObj, err := h.repos["imageRepostiory"].(*image.Repository).GetById(ctx, addImageRequest.ImageId)

	if err != nil {
		slog.Error(fmt.Sprintf("Unable to get image with : %s", addImageRequest.ImageId), "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Unable to find image with given id"})
		return
	}

	err = h.repos["groupRepository"].(*Repository).AddImage(ctx, groupId, imageObj.Id)

	if err != nil {
		slog.Error(fmt.Sprintf("Unable to add image with : %s to group: %s", addImageRequest.ImageId, groupId), "error", err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "Unable to add image to the group"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Image added successfully",
	})
}

func (h *Handler) addUser(ctx *gin.Context) {
	groupId := ctx.Param("id")
	addImageRequest := &AddImageRequest{}
	err := ctx.ShouldBindBodyWithJSON(addImageRequest)
	if err != nil {
		slog.Error("Error parsing the payload for add image request: ", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Please provide image id in request body"})
		return
	}

	userObj, err := h.repos["userRepository"].(*user.Repository).GetById(ctx, addImageRequest.ImageId)

	if err != nil {
		slog.Error(fmt.Sprintf("Unable to get user with : %s", addImageRequest.ImageId), "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Unable to find user with given id"})
		return
	}

	err = h.repos["groupRepository"].(*Repository).AddUser(ctx, groupId, userObj.Id)

	if err != nil {
		slog.Error(fmt.Sprintf("Unable to get image with : %s", addImageRequest.ImageId), "error", err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "Unable to add user to the group"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User added successfully",
	})
}

func (h *Handler) getImages(ctx *gin.Context) {
	groupId := ctx.Param("id")
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

	imageIds, err := h.repos["groupRepository"].(*Repository).GetImageIds(ctx, groupId, pageInt, limitInt)
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to get images for groupId : %s", groupId), "error", err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "Unable to images for the group"})
		return
	}

	images, err := h.repos["imageRepository"].(*image.Repository).FindByIds(ctx, imageIds)
	getImageResponseArr := []image.GetImageResponse{}
	for _, imageObj := range images {
		var getImageResponse image.GetImageResponse
		getImageResponse.Set(ctx, h.cloudFrontUrl, &imageObj)
		getImageResponseArr = append(getImageResponseArr, getImageResponse)
	}
	ctx.JSON(http.StatusOK, getImageResponseArr)

}

func (h *Handler) getUsers(ctx *gin.Context) {
	groupId := ctx.Param("id")
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

	userIds, err := h.repos["groupRepository"].(*Repository).GetUserIds(ctx, groupId, pageInt, limitInt)
	if err != nil {
		slog.Error(fmt.Sprintf("Unable to get users for groupId : %s", groupId), "error", err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "Unable to users for the group"})
		return
	}

	users, err := h.repos["userRepository"].(*user.Repository).FindByIds(ctx, userIds)
	getUserResponseArr := []user.GetUserResponse{}
	for _, userObj := range users {
		var getUserResponse user.GetUserResponse
		getUserResponse.Set(&userObj)
		getUserResponseArr = append(getUserResponseArr, getUserResponse)
	}
	ctx.JSON(http.StatusOK, getUserResponseArr)
}

func (h *Handler) deleteImage(ctx *gin.Context) {

}

func (h *Handler) deleteUser(ctx *gin.Context) {

}
