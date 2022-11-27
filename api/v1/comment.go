package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog-app/api/models"
	"github.com/ibrat-muslim/blog-app/storage/repo"
)

// @Router /comments [post]
// @Summary Create a comment
// @Description Create a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body models.CreateCommentRequest true "Comment"
// @Success 201 {object} models.Comment
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateComment(ctx *gin.Context) {

	var req models.CreateCommentRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Comment().Create(&repo.Comment{
		PostID:      req.PostID,
		UserID:      req.UserID,
		Description: req.Description,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Comment{
		ID:          resp.ID,
		PostID:      resp.PostID,
		UserID:      resp.UserID,
		Description: resp.Description,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
	})
}

// @Router /comments/{id} [get]
// @Summary Get a comment by id
// @Description Get a comment by id
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Comment
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetComment(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Comment().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Comment{
		ID:          resp.ID,
		PostID:      resp.PostID,
		UserID:      resp.UserID,
		Description: resp.Description,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
	})
}

func validateGetCommentsParams(ctx *gin.Context) (*repo.GetCommentsParams, error) {
	var (
		limit int64 = 10
		page  int64 = 1
		err   error
	)

	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &repo.GetCommentsParams{
		Limit: int32(limit),
		Page:  int32(page),
	}, nil
}

// @Router /comments [get]
// @Summary Get comments
// @Description Get comments
// @Tags comment
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success 200 {object} models.GetCommentsResult
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetComments(ctx *gin.Context) {
	queryParams, err := validateGetCommentsParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Comment().GetAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Router /comments/{id} [put]
// @Summary Update a comment
// @Description Update a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param comment body models.CreateCommentRequest true "Comment"
// @Success 200 {object} models.OKResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateComment(ctx *gin.Context) {
	var req models.CreateCommentRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	updatedAt := time.Now()

	err = h.storage.Comment().Update(&repo.Comment{
		ID:          id,
		PostID:      req.PostID,
		UserID:      req.UserID,
		Description: req.Description,
		UpdatedAt:   &updatedAt,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.OKResponse{
		Success: "successfully updated",
	})
}

// @Router /comments/{id} [delete]
// @Summary Delete a comment
// @Description Delete a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.OKResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteComment(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.storage.Comment().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.OKResponse{
		Success: "successfully deleted",
	})
}