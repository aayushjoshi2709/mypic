package group

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repos map[string]any
}

func (h *Handler) New(repos map[string]any) {
	h.repos = repos
}

func (h *Handler) add(ctx *gin.Context) {

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
	page := ctx.Param("page")
	limit := ctx.Param("limit")
	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		slog.Error("Error converting page variable", "error", err)
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "The value provided for page is not valid"})
		return
	}

	if limit == "" {
		limit = "10"
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

}

func (h *Handler) addUser(ctx *gin.Context) {

}

func (h *Handler) getImages(ctx *gin.Context) {

}

func (h *Handler) getUsers(ctx *gin.Context) {

}

func (h *Handler) deleteImage(ctx *gin.Context) {

}

func (h *Handler) deleteUser(ctx *gin.Context) {

}
