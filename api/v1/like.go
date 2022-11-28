package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog-app/api/models"
	"github.com/ibrat-muslim/blog-app/storage/repo"
)

// @Router /likes [post]
// @Summary Create like
// @Description Create like
// @Tags like
// @Accept json
// @Produce json
// @Param like body models.CreateLikeRequest true "Like"
// @Success 201 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateLike(ctx *gin.Context) {

	var req models.CreateLikeRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Like().Create(&repo.Like{
		PostID: req.PostID,
		UserID: req.UserID,
		Status: req.Status,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Like{
		ID:     resp.ID,
		PostID: resp.PostID,
		UserID: resp.UserID,
		Status: resp.Status,
	})
}

// @Router /likes/user-post [get]
// @Summary Get like by user and post
// @Description Get like by user and post
// @Tags like
// @Accept json
// @Produce json
// @Param post_id query int true "Post ID"
// @Success 200 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetLike(ctx *gin.Context) {

	post_id, err := strconv.ParseInt(ctx.Param("post_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var user_id int64 = 10 //TODO

	resp, err := h.storage.Like().Get(user_id, post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Like{
		ID:     resp.ID,
		PostID: resp.PostID,
		UserID: resp.UserID,
		Status: resp.Status,
	})
}
