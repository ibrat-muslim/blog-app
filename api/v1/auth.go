package v1

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog-app/api/models"
	"github.com/ibrat-muslim/blog-app/pkg/utils"
	"github.com/ibrat-muslim/blog-app/storage/repo"
)

var ErrWrongEmailOrPass = errors.New("wrong email or password")

// @Router /auth/register [post]
// @Summary Register a user
// @Description Register a user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Data"
// @Success 201 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Register(ctx *gin.Context) {

	var req models.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	result, err := h.storage.User().Create(&repo.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		Username:  req.Username,
		Type:      repo.UserTypeUser,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token, _, err := utils.CreateToken(&utils.TokenParams{
		UserID:   result.ID,
		Username: result.Username,
		Email:    result.Email,
		Duration: time.Hour * 24,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.AuthResponse{
		ID:          result.ID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		Username:    result.Username,
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		AccessToken: token,
	})
}

// @Router /auth/login [post]
// @Summary Login user
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.LoginRequest true "Data"
// @Success 201 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Login(ctx *gin.Context) {

	var req models.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.User().GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusForbidden, errorResponse(ErrWrongEmailOrPass))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, result.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(ErrWrongEmailOrPass))
		return
	}

	token, _, err := utils.CreateToken(&utils.TokenParams{
		UserID:   result.ID,
		Username: result.Username,
		Email:    result.Email,
		Duration: time.Hour * 24,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.AuthResponse{
		ID:          result.ID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		Username:    result.Username,
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		AccessToken: token,
	})
}
