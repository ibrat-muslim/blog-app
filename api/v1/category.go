package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/blog-app/api/models"
	"github.com/ibrat-muslim/blog-app/storage/repo"
)

// @Router /categories [post]
// @Summary Create a category
// @Description Create a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryRequest true "Category"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCategory(ctx *gin.Context) {

	var req models.CreateCategoryRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().Create(&repo.Category{
		Title: req.Title,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Category{
		ID:        resp.ID,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}

// @Router /categories/{id} [get]
// @Summary Get a category by id
// @Description Get a category by id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCategory(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Category{
		ID:        resp.ID,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}

func validateGetCategoriesParams(ctx *gin.Context) (*repo.GetCategoriesParams, error) {
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

	return &repo.GetCategoriesParams{
		Limit:  int32(limit),
		Page:   int32(page),
		Title: ctx.Query("title"),
	}, nil
}

// @Router /categories [get]
// @Summary Get categories
// @Description Get categories
// @Tags category
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param title query string false "Title"
// @Success 200 {object} models.GetCategoriesResult
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCategories(ctx *gin.Context) {
	queryParams, err := validateGetCategoriesParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Category().GetAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Router /categories/{id} [put]
// @Summary Update a category
// @Description Update a category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param category body models.CreateCategoryRequest true "Category"
// @Success 200 {object} models.OKResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateCategory(ctx *gin.Context) {
	var req models.CreateCategoryRequest

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

	err = h.storage.Category().Update(&repo.Category{
		ID: id,
		Title: req.Title,
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

// @Router /categories/{id} [delete]
// @Summary Delete a category
// @Description Delete a category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.OKResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.storage.Category().Delete(id)
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