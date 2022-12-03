package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog-app/api/models"
	emailPkg "github.com/ibrat-muslim/blog-app/pkg/email"
	"github.com/ibrat-muslim/blog-app/pkg/utils"
	"github.com/ibrat-muslim/blog-app/storage/repo"
)

// @Router /auth/register [post]
// @Summary Register a user
// @Description Register a user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Data"
// @Success 201 {object} models.OKResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Register(ctx *gin.Context) {

	var req models.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = h.storage.User().GetByEmail(req.Email)
	if !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusBadRequest, errorResponse(ErrEmailExists))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user := &repo.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
		Type:      repo.UserTypeUser,
	}

	userData, err := json.Marshal(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = h.inMemory.Set("user_"+user.Email, string(userData), 10*time.Minute)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	go func() {
		err := h.sendVerificationCode(req.Email)
		if err != nil {
			fmt.Printf("failed to send verification code: %v", err)
		}
	}()

	ctx.JSON(http.StatusCreated, models.OKResponse{
		Message: "Verification code has been sent!",
	})
}

func (h *handlerV1) sendVerificationCode(email string) error {
	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		return err
	}

	err = h.inMemory.Set("code_"+email, code, time.Minute)
	if err != nil {
		return err
	}

	err = emailPkg.SendEmail(h.cfg, &emailPkg.SendEmailRequest{
		To:      []string{email},
		Subject: "Verification email",
		Body: map[string]string{
			"code": code,
		},
		Type: emailPkg.VerificationEmail,
	})
	if err != nil {
		return err
	}

	return nil
}

// @Router /auth/verify [post]
// @Summary Verify user
// @Description Verify user
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 201 {object} models.AuthResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Verfiy(ctx *gin.Context) {

	var req models.VerifyRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userData, err := h.inMemory.Get("user_" + req.Email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	var user repo.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	code, err := h.inMemory.Get("code_" + user.Email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(ErrCodeExpired))
		return
	}

	if req.Code != code {
		ctx.JSON(http.StatusForbidden, errorResponse(ErrIncorrectCode))
		return
	}

	result, err := h.storage.User().Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token, _, err := utils.CreateToken(&utils.TokenParams{
		UserID:   result.ID,
		UserType: result.Type,
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
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
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
		UserType: result.Type,
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
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		AccessToken: token,
	})
}
