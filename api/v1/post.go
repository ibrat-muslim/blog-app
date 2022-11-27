package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog-app/api/models"
	"github.com/ibrat-muslim/blog-app/storage/repo"
)

// @Router /posts [post]
// @Summary Create a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body models.CreatePostRequest true "Post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreatePost(ctx *gin.Context) {

	var req models.CreatePostRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Post().Create(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Post{
		ID:          resp.ID,
		Title:       resp.Title,
		Description: resp.Description,
		ImageUrl:    resp.ImageUrl,
		UserID:      resp.UserID,
		CategoryID:  resp.CategoryID,
		CreatedAt:   resp.CreatedAt,
	})
}

// @Router /posts/{id} [get]
// @Summary Get a post by id
// @Description Get a post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetPost(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Post().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Post{
		ID:          resp.ID,
		Title:       resp.Title,
		Description: resp.Description,
		ImageUrl:    resp.ImageUrl,
		UserID:      resp.UserID,
		CategoryID:  resp.CategoryID,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
		ViewsCount:  resp.ViewsCount,
	})
}
