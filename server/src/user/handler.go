package user

import (
	"log/slog"
	"strings"

	"net/http"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/aayushjoshi2709/mypic/src/utils/encrypt"
	"github.com/aayushjoshi2709/mypic/src/utils/jwt"
	"github.com/aayushjoshi2709/mypic/src/utils/redis"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repos           map[string]any
	loggedOutPrefix string
}

func (h *Handler) New(repos map[string]any) {
	h.repos = repos
	h.loggedOutPrefix = "logged_out_"
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a user by their unique ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/user/{id} [get]
func (h *Handler) get(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.repos["userRepository"].(*Repository).GetById(
		ctx,
		id,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while fetching the user"})
		slog.Error("Error fetching user", "error", err)
		return
	}

	var getUserResponse GetUserResponse
	getUserResponse.Set(user)
	ctx.JSON(http.StatusOK, getUserResponse)
}

func (h *Handler) getAll(ctx *gin.Context) {}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User details"
// @Success 201 {object} GetUserResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Router /api/v1/user [post]
func (h *Handler) create(ctx *gin.Context) {
	var createUserRequest CreateUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&createUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: err.Error()})
		return
	}

	user, err := h.repos["userRepository"].(*Repository).GetByUsername(ctx, createUserRequest.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while fetching the user"})
		slog.Error("Error fetching user", "error", err)
		return
	}

	if user != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "Username already exists"})
		return
	}

	hashedPassword, err := encrypt.GenerateFromPassword(createUserRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "An error occurred while hashing the password"})
		slog.Error("Error hashing password", "error", err)
		return
	}
	createUserRequest.Password = hashedPassword

	user, err = h.repos["userRepository"].(*Repository).Add(
		ctx,
		createUserRequest.Name,
		createUserRequest.Username,
		createUserRequest.Password,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while creating the user"})
		slog.Error("Error creating user", "error", err)
		return
	}

	var getUserResponse GetUserResponse
	getUserResponse.Set(user)
	ctx.JSON(http.StatusCreated, getUserResponse)
}

// @UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "Updated user details"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Failure 404 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/user/{id} [put]
func (h *Handler) update(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateUserRequest UpdateUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&updateUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: err.Error()})
		return
	}

	user, err := h.repos["userRepository"].(*Repository).Update(
		ctx,
		id,
		updateUserRequest.Name,
		updateUserRequest.Username,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while updating the user"})
		slog.Error("Error updating user", "error", err)
		return
	}

	var getUserResponse GetUserResponse
	getUserResponse.Set(user)
	ctx.JSON(http.StatusOK, getUserResponse)
}

// @DeleteUser godoc
// @Summary Delete an existing user
// @Description Delete an existing user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/user/{id} [delete]
func (h *Handler) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.repos["userRepository"].(*Repository).Delete(
		ctx,
		id,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while deleting the user"})
		slog.Error("Error deleting user", "error", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// @LoginUser godoc
// @Summary Login a user
// @Description Login a user with their username and password
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body LoginUserRequest true "User credentials"
// @Success 200 {object} LoginUserResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Failure 401 {object} common.ErrorResponseDto
// @Router /api/v1/user/login [post]
func (h *Handler) login(ctx *gin.Context) {
	var loginUserRequest LoginUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&loginUserRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: err.Error()})
		return
	}

	user, err := h.repos["userRepository"].(*Repository).GetByUsername(
		ctx,
		loginUserRequest.Username,
	)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, common.ErrorResponseDto{Error: "Invalid username"})
		slog.Error("Error fetching user by username", "error", err)
		return
	}

	bcryptEncodedPassword := user.Password
	err = encrypt.CompareHashAndPassword(bcryptEncodedPassword, loginUserRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, common.ErrorResponseDto{Error: "Invalid password"})
		slog.Error("Error comparing passwords", "error", err)
		return
	}

	token, err := jwt.Init().GenerateToken(user.Username, user.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponseDto{Error: "An error occurred while generating the token"})
		slog.Error("Error generating token", "error", err)
		return
	}

	var loginUserResponse LoginUserResponse
	loginUserResponse.Token = token
	ctx.JSON(http.StatusOK, loginUserResponse)
}

// @LogoutUser godoc
// @Summary Logout a user
// @Description Logout a user by invalidating their token
// @Tags Users
// @Accept json
// @Produce json
// @Param token body LogoutUserRequest true "Logout payload"
// @Success 200 {object} LogoutUserResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/user/logout [delete]
func (h *Handler) logout(ctx *gin.Context) {
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

	tokenString, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in context"})
		ctx.Abort()
		return
	}
	redisKey := h.loggedOutPrefix + tokenString.(string)
	redis.Init().Set(ctx, redisKey, "invalid", jwt.GetExprityInDays())
	ctx.JSON(http.StatusOK, LogoutUserResponse{Message: "Successfully logged out"})
}

// GetCurrentUser godoc
// @Summary Get current user details
// @Description Get current user details
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} common.ErrorResponseDto
// @Security BearerAuth
// @Router /api/v1/user/me [get]
func (h *Handler) getCurrentUser(ctx *gin.Context) {
	user, err := h.repos["userRepository"].(*Repository).getCurrentUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponseDto{Error: "An error occurred while fetching the user"})
		slog.Error("Error fetching user by ID", "error", err)
		return
	}

	var getUserResponse GetUserResponse
	getUserResponse.Set(user)
	ctx.JSON(http.StatusOK, getUserResponse)
}
