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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post := parsePostToModel(resp)

	ctx.JSON(http.StatusCreated, post)
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Post().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post := parsePostToModel(resp)

	likeInfo, err := h.storage.Like().GetLikesDislikesCount(post.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post.LikeInfo = &models.PostLikeInfo{
		LikesCount:    likeInfo.LikesCount,
		DislikesCount: likeInfo.DislikesCount,
	}

	ctx.JSON(http.StatusOK, post)
}

// @Router /posts [get]
// @Summary Get posts
// @Description Get posts
// @Tags post
// @Accept json
// @Produce json
// @Param filter query models.GetAllParamsRequest false "Filter"
// @Success 200 {object} models.GetPostsResult
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetPosts(ctx *gin.Context) {
	request, err := validateGetAllParamsRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Post().GetAll(&repo.GetPostsParams{
		Limit:  request.Limit,
		Page:   request.Page,
		Search: request.Search,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getPostsResponse(result))
}

func getPostsResponse(data *repo.GetPostsResult) *models.GetPostsResponse {
	response := models.GetPostsResponse{
		Posts: make([]*models.Post, 0),
		Count: data.Count,
	}

	for _, post := range data.Posts {
		p := parsePostToModel(post)
		response.Posts = append(response.Posts, &p)
	}

	return &response
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updatedAt := time.Now()

	err = h.storage.Post().Update(&repo.Post{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
		UpdatedAt:   &updatedAt,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.storage.Post().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.OKResponse{
		Success: "successfully deleted",
	})
}

func parsePostToModel(post *repo.Post) models.Post {
	return models.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserID:      post.UserID,
		CategoryID:  post.CategoryID,
		CreatedAt:   post.CreatedAt,
		ViewsCount:  post.ViewsCount,
	}
}