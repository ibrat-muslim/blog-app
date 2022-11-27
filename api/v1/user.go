package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog-app/api/models"
	"github.com/ibrat-muslim/blog-app/storage/repo"
)

// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User"
// @Success 201 {object} models.User
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateUser(ctx *gin.Context) {

	var req models.CreateUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.User().Create(&repo.User{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		PhoneNumber:     req.PhoneNumber,
		Email:           req.Email,
		Gender:          req.Gender,
		Password:        req.Password,
		Username:        req.Username,
		ProfileImageUrl: req.ProfileImageUrl,
		Type:            req.Type,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.User{
		ID:              resp.ID,
		FirstName:       resp.FirstName,
		LastName:        resp.LastName,
		PhoneNumber:     resp.PhoneNumber,
		Email:           resp.Email,
		Gender:          resp.Gender,
		Username:        resp.Username,
		ProfileImageUrl: resp.ProfileImageUrl,
		Type:            resp.Type,
		CreatedAt:       resp.CreatedAt,
	})
}

// @Router /users/{id} [get]
// @Summary Get a user by id
// @Description Get a user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.User
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetUser(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.User().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.User{
		ID:              resp.ID,
		FirstName:       resp.FirstName,
		LastName:        resp.LastName,
		PhoneNumber:     resp.PhoneNumber,
		Email:           resp.Email,
		Gender:          resp.Gender,
		Username:        resp.Username,
		ProfileImageUrl: resp.ProfileImageUrl,
		Type:            resp.Type,
		CreatedAt:       resp.CreatedAt,
	})
}
