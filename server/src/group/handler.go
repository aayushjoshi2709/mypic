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

// @Summary Add new group
// @Description  Add new group
// @Tags Group
// @Accept json
// @Produce json
// @Param group body CreateGroupRequest true "Group details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group [post]
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

// @Summary Get an group by ID
// @Description Get an group by its unique ID
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} GetGroupResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group/{id} [get]
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

// @Summary Get all groups
// @Description Get all group
// @Tags Group
// @Accept json
// @Produce json
// @Param page query number false "Page number" default(1)
// @Param limit query number false "Number of groups per page" default(10)
// @Success 200 {object} []GetGroupResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group [get]
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

// @Summary Delete a group
// @Description Delete a group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 204 "No Content"
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group/{id} [delete]
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

// @Summary Add image to the group
// @Description Add image to the group
// @Tags Group
// @Accept json
// @Produce json
// @Param image body AddImageRequest true "Image details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group/{id}/image [post]
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

// @Summary Add user to the group
// @Description Add user to the group
// @Tags Group
// @Accept json
// @Produce json
// @Param user body AddUserRequest true "User details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group/{id}/user [post]
func (h *Handler) addUser(ctx *gin.Context) {
	groupId := ctx.Param("id")
	addUserRequest := &AddUserRequest{}
	err := ctx.ShouldBindBodyWithJSON(addUserRequest)
	if err != nil {
		slog.Error("Error parsing the payload for add user request: ", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Please provide user id in request body"})
		return
	}

	userObj, err := h.repos["userRepository"].(*user.Repository).GetById(ctx, addUserRequest.UserId)

	if err != nil {
		slog.Error(fmt.Sprintf("Unable to get user with : %s", addUserRequest.UserId), "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Unable to find user with given id"})
		return
	}

	err = h.repos["groupRepository"].(*Repository).AddUser(ctx, groupId, userObj.Id)

	if err != nil {
		slog.Error(fmt.Sprintf("Unable to get user with : %s", addUserRequest.UserId), "error", err)
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "Unable to add user to the group"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User added successfully",
	})
}

// @Summary Get all images of the group
// @Description Get all images of the group
// @Tags Group
// @Accept json
// @Produce json
// @Param page query number false "Page number" default(1)
// @Param limit query number false "Number of images per page" default(10)
// @Success 200 {object} []image.GetImageResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group/{id}/image [get]
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

// @Summary Get all users of the group
// @Description Get all users of the group
// @Tags Group
// @Accept json
// @Produce json
// @Param page query number false "Page number" default(1)
// @Param limit query number false "Number of users per page" default(10)
// @Success 200 {object} []user.GetUserResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group/{id}/user [get]
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

// @Summary Delete a image from the group
// @Description Delete a image from the group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param imageId path string true "Image ID"
// @Success 204 "No Content"
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group/{id}/image/{imageId} [delete]
func (h *Handler) deleteImage(ctx *gin.Context) {
	groupId := ctx.Param("id")
	imageId := ctx.Param("imageId")
	err := h.repos["groupRepository"].(*Repository).RemoveImage(ctx, groupId, imageId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Error deleting image id from group"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary Delete a user from the group
// @Description Delete a user from the group
// @Tags Group
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param userId path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/group/{id}/user/{userId} [delete]
func (h *Handler) deleteUser(ctx *gin.Context) {
	groupId := ctx.Param("id")
	userId := ctx.Param("userId")
	err := h.repos["groupRepository"].(*Repository).RemoveUser(ctx, groupId, userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Error deleting user id from group"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
