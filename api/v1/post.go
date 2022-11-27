package v1

import (
	"net/http"
	"strconv"
	"time"

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

func validateGetPostsParams(ctx *gin.Context) (*repo.GetPostsParams, error) {
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

	return &repo.GetPostsParams{
		Limit:  int32(limit),
		Page:   int32(page),
		Title: ctx.Query("title"),
	}, nil
}

// @Router /posts [get]
// @Summary Get posts
// @Description Get posts
// @Tags post
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param title query string false "Title"
// @Success 200 {object} models.GetPostsResult
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetPosts(ctx *gin.Context) {
	queryParams, err := validateGetPostsParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Post().GetAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Router /posts/{id} [put]
// @Summary Update a post
// @Description Update a post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param post body models.CreatePostRequest true "Post"
// @Success 200 {object} models.OKResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdatePost(ctx *gin.Context) {
	var req models.CreatePostRequest

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

	err = h.storage.Post().Update(&repo.Post{
		ID: id,
		Title: req.Title,
		Description: req.Description,
		ImageUrl: req.ImageUrl,
		UserID: req.UserID,
		CategoryID: req.CategoryID,
		UpdatedAt: &updatedAt,
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

// @Router /posts/{id} [delete]
// @Summary Delete a post
// @Description Delete a post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.OKResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeletePost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.storage.Post().Delete(id)
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