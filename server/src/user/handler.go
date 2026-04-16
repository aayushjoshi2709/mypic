package user

import (
	"log/slog"

	"github.com/aayushjoshi2709/mypic/src/common"
	"github.com/aayushjoshi2709/mypic/src/utils/encrypt"
	"github.com/aayushjoshi2709/mypic/src/utils/jwt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func (h *Handler) New(repo *Repository) {
	h.repo = repo
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
// @Router /api/v1/user/{id} [get]
func (h *Handler) get(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.repo.GetById(
		ctx.Request.Context(),
		id,
	)

	if err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "An error occurred while fetching the user"})
		slog.Error("Error fetching user", "error", err)
		return
	}

	var getUserResponse GetUserResponse
	getUserResponse.Set(user)
	ctx.JSON(200, getUserResponse)
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
		ctx.JSON(400, common.ErrorResponseDto{Error: err.Error()})
		return
	}

	hashedPassword, err := encrypt.GenerateFromPassword(createUserRequest.Password)
	if err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "An error occurred while hashing the password"})
		slog.Error("Error hashing password", "error", err)
		return
	}
	createUserRequest.Password = hashedPassword

	user, err := h.repo.Add(
		ctx.Request.Context(),
		createUserRequest.Name,
		createUserRequest.Username,
		createUserRequest.Password,
	)

	if err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "An error occurred while creating the user"})
		slog.Error("Error creating user", "error", err)
		return
	}

	var getUserResponse GetUserResponse
	getUserResponse.Set(user)
	ctx.JSON(201, getUserResponse)
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
// @Router /api/v1/user/{id} [put]
func (h *Handler) update(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateUserRequest UpdateUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&updateUserRequest); err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: err.Error()})
		return
	}

	user, err := h.repo.Update(
		ctx.Request.Context(),
		id,
		updateUserRequest.Name,
		updateUserRequest.Username,
	)
	if err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "An error occurred while updating the user"})
		slog.Error("Error updating user", "error", err)
		return
	}

	var getUserResponse GetUserResponse
	getUserResponse.Set(user)
	ctx.JSON(200, getUserResponse)
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
// @Router /api/v1/user/{id} [delete]
func (h *Handler) delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.repo.Delete(
		ctx.Request.Context(),
		id,
	)

	if err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "An error occurred while deleting the user"})
		slog.Error("Error deleting user", "error", err)
		return
	}

	ctx.Status(204)
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
// @Router /api/v1/user/login [post]
func (h *Handler) login(ctx *gin.Context) {
	var loginUserRequest LoginUserRequest
	if err := ctx.ShouldBindBodyWithJSON(&loginUserRequest); err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: err.Error()})
		return
	}

	user, err := h.repo.GetByUsername(
		ctx.Request.Context(),
		loginUserRequest.Username,
	)

	if err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "Invalid username"})
		slog.Error("Error fetching user by username", "error", err)
		return
	}

	bcryptEncodedPassword := user.Password
	err = encrypt.CompareHashAndPassword(bcryptEncodedPassword, loginUserRequest.Password)
	if err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "Invalid password"})
		slog.Error("Error comparing passwords", "error", err)
		return
	}

	token, err := jwt.GenerateToken(user.Username, user.Id)
	
	if err != nil {
		ctx.JSON(400, common.ErrorResponseDto{Error: "An error occurred while generating the token"})
		slog.Error("Error generating token", "error", err)	
	}
	
	var loginUserResponse LoginUserResponse
	loginUserResponse.Token = token
	ctx.JSON(200, loginUserResponse)
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
// @Router /api/v1/user/logout [post]
func (h *Handler) logout(ctx *gin.Context) {}
